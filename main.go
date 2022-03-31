package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matumoto1234/go-todo/chat"
)

func ping_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	handler := chat.TemplateHandler{Filename: "chat.html"}
	http.Handle("/", &handler)

	http.HandleFunc("/ping", ping_handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
