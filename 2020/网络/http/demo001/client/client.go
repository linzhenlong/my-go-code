package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type myRead struct {
}

func (m *myRead) Read(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return 0, nil
}

func main() {
	resp, err := http.Get("http://127.0.0.1:6060/go")
	/* if err != nil {
		panic(err)
	} */
	defer func() {

		r := recover()
		if r != nil {
			fmt.Println(err)
			return
		}
		resp.Body.Close()
	}()
	bytes, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
