package main

import (
	"fmt"
	"os"
	"text/template"
)

type Preson struct {
	Name string
	Age int
}

func main() {

	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("template.ParseFiles error =", err)
		return
	}
	p := Preson{
		Name:"TOM",
		Age:18,
	}
	if err:= t.Execute(os.Stdout,p);err != nil {
		fmt.Println("t.Execute(os.Stdout,p) error=", err)
		return
	}

}
