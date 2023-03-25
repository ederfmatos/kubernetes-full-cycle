package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/health", Health)
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

func Health(writer http.ResponseWriter, request *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 10 || duration.Seconds() > 30 {
		writer.WriteHeader(500)
		writer.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		writer.WriteHeader(200)
		writer.Write([]byte("ok"))
	}
}
