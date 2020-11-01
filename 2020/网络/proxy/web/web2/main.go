package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("this is web 2 9092")
		//time.Sleep(time.Second * 3)
		w.Write([]byte("web2"))
	})
	http.ListenAndServe(":9092", nil)
}
