package main

import "net/http"

import "io"

func sendResponse(w http.ResponseWriter,sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}