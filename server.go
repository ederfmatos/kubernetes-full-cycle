package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/secrets", Secrets)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(writer, "Hello, I'm %s. I'm %s.", name, age)
}

func Secrets(writer http.ResponseWriter, request *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(writer, "User %s. Password %s.", user, password)
}
