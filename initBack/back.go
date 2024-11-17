package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("usernameis:%s. heis %d and he has money %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"BOB", 25, -50, 4.2, 0.8}
	bob.setNewName("alex")
	fmt.Fprintf(w, bob.getAllInfo())
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ContactsPage")
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {

	HandleRequest()
}
