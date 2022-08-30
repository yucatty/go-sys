package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter implements Write([]byte) (int, error). Therefore you can use it everywhere where a io.Writer is required.
	io.WriteString(w, "http.ResponseWriter sample")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
