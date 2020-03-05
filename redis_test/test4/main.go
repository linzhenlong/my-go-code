package main

import(
	"net/http"
	"log"
	"io/ioutil"
	"os"

)

type appHandler func(w http.ResponseWriter, r *http.Request) error

// 捕获异常
func catchPanic(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			log.Printf("catch panic: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func list(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	contents , err := ioutil.ReadAll(file)
	if err !=nil {
		return err
	}
	w.Write(contents)
	return nil
}
func main() {
	http.HandleFunc("/list/", catchPanic(list))
	http.ListenAndServe(":9000", nil)
}