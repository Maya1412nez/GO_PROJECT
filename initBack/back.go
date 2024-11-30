package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ContactsPage")
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/main.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "main.html", nil)
}

func HandleRequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {

	HandleRequest()
}
