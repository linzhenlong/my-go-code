package util

import "time"

import "net/http"

// HTTPChecker ...
type HTTPChecker struct {
	Servers HTTPServers
}

// NewHTTPChecker 实例化.
func NewHTTPChecker(servers HTTPServers) *HTTPChecker {
	return &HTTPChecker{
		Servers: servers,
	}
}

// Check 检测..
func (h *HTTPChecker) Check(timeOut time.Duration) {
	httpClient := http.Client{}
	for _, webServer := range h.Servers {
		// head 请求 目标网站
		resp, err := httpClient.Head(webServer.Host)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			webServer.Status = "DOWN"
			continue
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			webServer.Status = "UP"
		} else {
			webServer.Status = "DOWN"
		}
	}
}
