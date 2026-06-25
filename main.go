package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Go JSON API is running")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/student", createStudent)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}