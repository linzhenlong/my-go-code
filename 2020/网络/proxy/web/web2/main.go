package main

import "net/http"

import "log"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("this is web 2 9092")
		w.Write([]byte("web2"))
	})
	http.ListenAndServe(":9092", nil)
}
