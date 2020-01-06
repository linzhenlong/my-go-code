package main

import (
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/dbops"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/session"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/defs"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"io"
	"net/http"
)
// CreateUser 创建用户.
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	//io.WriteString(w, "Create user handler")
	res , _ := ioutil.ReadAll(r.Body)
	MyLog.Infof("r==>", r)
	r.ParseForm()
	MyLog.Debug("PostFormValue user_name",r.Form["user_name"])
	MyLog.Debug("PostFormValue pwd",r.PostFormValue("pwd"))
	MyLog.Info("CreateUser res=", string(res))
	uBody := &defs.UserCredential{}
	MyLog.Info("uBody=", uBody)
	// 这块普通的http Post HEAD 是Content-Type:[multipart/form-data json解析有问题.
	if err := json.Unmarshal(res, &uBody); err !=nil {
		MyLog.Info("json.Unmarshal(res, uBody)", uBody)
		MyLog.Info("json.Unmarshal(res, uBody) error", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	
	if err := dbops.AddUserCredential(uBody.Username, uBody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	sessionID := session.GenerateNewSessionID(uBody.Username)
	signedUP := &defs.SignedUp{}
	signedUP.SessionID = sessionID
	signedUP.Suceess = true
	if resp, err := json.Marshal(signedUP);err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)	
	} else {
		sendNormalResponse(w, string(resp),201)
	}
	return

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("user_name")

	MyLog.Info(p)
	MyLog.Info(uname)
	io.WriteString(w, uname)
}