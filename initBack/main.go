package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ti pidor")
	t, err := template.ParseFiles("templates/main.html")
	if err != nil {
		fmt.Println("cant'parse files main.html")
	}
	t.Execute(w, "main.html")
	// return nil
}

func HandleRequest() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":5000", nil)
}

func main() {
	fmt.Printf("initializing back on :5000")
	HandleRequest()
}
