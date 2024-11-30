package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ti pidor")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("cant'parse files index.html")
	}
	t.Execute(w, "index.html")
}

func HandleRequest() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func main() {
	fmt.Printf("initializing back on :5000")
	HandleRequest()
}
