package main

import (
	"net/http"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/session"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/defs"
	
)

// 自定义http header.
var headerFieldSession = "X-SESSION-ID"
var headerFieldUName = "X-User-Name"

// 校验session.
func validateUserSession(r *http.Request) bool {
	sessionID := r.Header.Get(headerFieldSession)
	// 如果session_id为空，返回false
	if len(sessionID) == 0 {
		return false
	}
	userName , ok := session.IsSessionExpired(sessionID)
	if ok {
		return false
	}
	r.Header.Add(headerFieldUName, userName)
	return true
}

// 校验用户.
func validateUser(w http.ResponseWriter, r *http.Request) bool{
	userName := r.Header.Get(headerFieldUName)
	if len(userName) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
