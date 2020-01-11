package taskrunner

import "github.com/linzhenlong/my-go-code/liumeiti/sys/scheduler/dbops"

import "log"

import "errors"

import "os"

import "sync"

// VideoClearDispatcher 视频清理.
func VideoClearDispatcher(dc dataChan) (err error) {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("VideoClearDispatcher error:=%s", err.Error())
		return 
	}
	if len(res) == 0 {
		return errors.New("no has delele videos")
	}
	for _, id := range res {
		dc <- id
	}
	return 
}

// DeleteVideo 删除视频.
func DeleteVideo(vid string)(err error) {
	err = os.Remove(videoPath + vid)
	// 
	if err !=nil && !os.IsNotExist(err) {
		log.Printf("DeleteVideo error:%s", err.Error())
		return
	}
	return nil
}

// VideoClearExecutor 执行.
func VideoClearExecutor(dc dataChan)(err error){
	errMap := &sync.Map{}
	LABLE:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := DeleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err!=nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
			default:
				break LABLE
		}
	}
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err !=nil {
			return false
		}
		return true
	})
	return nil
}