package main

import "net/http"

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>Hello Full Cycle</h1>"))
}
