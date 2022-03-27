package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! I am matumoto!!!!!!!!!!!!!!")
}

func ping_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/", handler)
  http.HandleFunc("/ping", ping_handler)
	http.ListenAndServe(":8080", nil)
}
