package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server is running on port 8080")

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `Welcome to my first golang webpage`)
}
func about(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `Welcome to my first about webpage`)
}
func contact(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `Welcome to my first contact webpage`)
}
