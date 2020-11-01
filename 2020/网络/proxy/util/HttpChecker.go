package util

import (
	"log"
	"net/http"
	"time"
)

// HTTPChecker ...
type HTTPChecker struct {
	Servers    HTTPServers
	FailMax    int // 失败的最大次数，达到这个值会被标识为Down
	RecovCount int // 连续成功,达到这个值，就会标识位UP
}

// NewHTTPChecker 实例化.
func NewHTTPChecker(servers HTTPServers) *HTTPChecker {
	return &HTTPChecker{
		Servers:    servers,
		FailMax:    6,
		RecovCount: 3,
	}
}

// Check 检测..
func (h *HTTPChecker) Check(timeOut time.Duration) {
	httpClient := http.Client{}
	httpClient.Timeout = timeOut

	for _, webServer := range h.Servers {
		// head 请求 目标网站
		resp, err := httpClient.Head(webServer.Host)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			log.Println(err)
			h.Fail(webServer)
			continue
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			h.Success(webServer)
		} else {
			h.Fail(webServer)
		}
	}
}

// Fail 失败的计数器...
func (h *HTTPChecker) Fail(server *HTTPServer) {
	if server.FailCount >= h.FailMax { // 如果失败的次数超过阈值
		server.Status = "DOWN"
	} else {
		server.FailCount++
	}
	server.SuccessCount = 0 // 必须连续成功，有一次失败就会置为0
}

// Success 成功的计数器.
func (h *HTTPChecker) Success(server *HTTPServer) {
	if server.FailCount > 0 {
		server.FailCount--
		server.SuccessCount++
		if server.SuccessCount == h.RecovCount {
			// 将失败次数置为可用节点
			server.FailCount = 0
			server.Status = "UP"
			server.SuccessCount = 0 // 重新计数
		}
	} else {
		server.Status = "UP"
	}
}
