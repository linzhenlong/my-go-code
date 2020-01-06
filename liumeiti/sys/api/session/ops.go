package session

import (
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/utils"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/defs"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/dbops"
	"time"
	"sync"
)


// 线程安全的map
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

// LoadSessionsFromDB 从数据库或许session信息.


func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key,ss)
		return true
	})
}

// GenerateNewSessionID 生成sessionID.
func GenerateNewSessionID(UserName string)(sessionID string) {
	sessionID, _ = utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000 // 过期时间半小时,*1000转毫秒.
	ss := &defs.SimpleSession{
		UserName: UserName,
		TTL: ttl,
	}

	// 放入缓存中
	sessionMap.Store(sessionID, ss)
	// 在放入数据库
	err := dbops.InsertSession(sessionID,ttl, UserName)
	if err != nil {
		return 
	}
	return sessionID
}

// IsSessionExpired 判断session是否过期.
func IsSessionExpired(sessionID string)(UserName string, status bool) {
	ss, ok := sessionMap.Load(sessionID)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			// 过期了，需要删除该session
			DeleteExpiredSession(sessionID)
			return "", true
		}
		return ss.(*defs.SimpleSession).UserName, false
	}
	return "", true

}

// nowInMilli 毫秒时间戳.
func nowInMilli() int64 {
	return  time.Now().UnixNano() / 100000 // 毫秒时间戳.
}

func DeleteExpiredSession(sessionId string) {
	sessionMap.Delete(sessionId)
	dbops.DeleteSession(sessionId)

}