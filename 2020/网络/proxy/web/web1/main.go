package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("this is web1")
		w.Write([]byte("web1"))
	})
	http.ListenAndServe(":9091", nil)
}
