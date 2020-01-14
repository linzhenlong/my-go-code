package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/scheduler/dbops"
	"net/http"
)

func videoDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, 400, "vid-id不能为空")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid) 
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, "MYSQL 写入失败"+err.Error())
		return
	}
	sendResponse(w, http.StatusOK, "success")

}