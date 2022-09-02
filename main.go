package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)

}

func HelloServer(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello server"))

}
