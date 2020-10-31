package main

import "net/http"

import "log"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("this is web3")
		w.Write([]byte(`web3`))
	})
	http.ListenAndServe(":9093", nil)
}
