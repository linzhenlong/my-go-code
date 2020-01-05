package main

import "net/http"

import "math/rand"

import "time"

import "fmt"

import "strconv"

func indexHandler(w http.ResponseWriter,r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(2)
	if randNum == 0 {
		time.Sleep(time.Second * 3)
		fmt.Fprintf(w,"slow response"+strconv.Itoa(randNum))
		return
	}
	fmt.Fprintf(w,"quick response"+strconv.Itoa(randNum))
}

func main() {
	http.HandleFunc("/",indexHandler)
	err := http.ListenAndServe("0.0.0.0:8889",nil)
	if err !=nil {
		panic(err)
	}
}