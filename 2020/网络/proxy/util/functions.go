package util

import (
	"net/http"
)

// CloneHeader 拷贝http header
func CloneHeader(src http.Header, dest *http.Header){
	for k, v:= range src {
		dest.Set(k,v[0])
	}
}