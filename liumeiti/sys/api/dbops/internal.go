package dbops

import (
	"database/sql"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/defs"
	"strconv"
	"sync"
)

// InsertSession 写入session.
func InsertSession(sessionID string, ttl int64, userName string) (err error) {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions(session_id,ttl,login_name) values(?, ?, ?)")
	if err != nil {
		return
	}
	_, err = stmtIns.Exec(sessionID, ttlStr, userName)
	if err != nil {
		return
	}
	defer stmtIns.Close()
	return
}

// RetrieveSession 获取session.
func RetrieveSession(sessionID string) (session *defs.SimpleSession, err error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select ttl,login_name from sessions where session_id=?")
	if err != nil {
		return ss, err
	}
	var (
		ttl       string
		loginName string
	)
	err = stmtOut.QueryRow(sessionID).Scan(&ttl, &loginName)

	if err != nil && err != sql.ErrNoRows {
		return
	}
	if err == sql.ErrNoRows {
		return
	}

	ttlInt, err := strconv.ParseInt(ttl, 10, 64)
	if err != nil {
		return nil, err
	}
	ss.UserName = loginName
	ss.TTL = int64(ttlInt)
	defer stmtOut.Close()
	return ss, nil
}

// RetrieveAllSessions 获取所有sessions.
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	sqlTemplate := "select session_id, ttl, login_name from sessions"
	stmt, err := dbConn.Prepare(sqlTemplate)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			ttl       string
			loginName string
			sessionID string
		)
		scanErr := rows.Scan(&sessionID, &ttl, &loginName)
		if scanErr != nil {
			break
		}
		ttlInt, _ := strconv.ParseInt(ttl, 10, 64)
		m.Store(sessionID, &defs.SimpleSession{
			UserName: loginName,
			TTL:      ttlInt,
		})
	}
	defer stmt.Close()
	return m, nil
}

// DeleteSession 删除session 信息.
func DeleteSession(sessionID string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM sessions where session_id=? limit 1")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(sessionID)
	if err != nil {
		return err
	}
	return nil
}
