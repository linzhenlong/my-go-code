package fetcher

import (
	"bufio"
	"fmt"
	"log"
	_"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	_"io"
	"io/ioutil"
	"net/http"
)

// Fetch 抓紧url内容.
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Response StatusCode is %d is not 200,%s", resp.StatusCode, resp.Status)
		return nil, err
	}

	bodyReader := bufio.NewReader(resp.Body)
	
	enCoding := determineEncoding(bodyReader)
	// gbk 转utf8
	utf8Reader := transform.NewReader(bodyReader, enCoding.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	//all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}

// determineEncoding 获取HTML-charset.
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil { // 如果Peek 不出来,就返回默认的utf-8编码.
		log.Printf("Fetcher error %v:",err)
		return unicode.UTF8
	}
	enCoding, _, _ := charset.DetermineEncoding(bytes, "")
	return enCoding
}
