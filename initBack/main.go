package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ti gnom...")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("cant'parse files index.html")
	}
	t.Execute(w, "index.html")
}

func HandleRequest() {
	http.HandleFunc("/", index)
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("../imgs"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.ListenAndServe(":9000", nil)
}

func main() {
	fmt.Printf("initializing back on :9000")
	HandleRequest()
}
