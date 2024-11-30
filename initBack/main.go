package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/main.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "main.html", nil)
}

func HandleRequest() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:5000", nil)
}

func main() {
	fmt.Printf("initializing back on 127.0.0.1:5000")
	HandleRequest()
}
