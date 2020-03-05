package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"os"
	"errors"
	"fmt"
)

func errPanic(w http.ResponseWriter, r *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}
func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(w http.ResponseWriter, r *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(w http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}
func errNotPermisson(w http.ResponseWriter, r *http.Request) error {
	return os.ErrPermission
}
func errUnkown(w http.ResponseWriter, r *http.Request) error {
	return errors.New("unkown error")
}
func noError(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "no error")
	return nil
}

// 测试数据
var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "服务器开小差了..."},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermisson, 403, "Forbidden"},
	{errUnkown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 通过httptest.NewRecorder() 及httptest.NewRequest(）
// 模拟http的请求
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com",
			nil,
		)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

// 实际起一个http server.
func TestErrWrapperINServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("epect(%d, %s);got (%d, %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
