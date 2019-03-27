package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/about'>About<h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/'>Index<h1>")
}

func main() {
	// Creating handlers
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server starting...")

	// Creating the server
	http.ListenAndServe(":3001", nil)
}
