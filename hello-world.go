package main

import (
	"fmt"
//	"hello-world/health"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello ,Welcome to the GOlang Application</h1>")
}

func main() {
	fmt.Println("Server Starting...")
	mux := http.NewServeMux()
	mux.HandleFunc("/new", index)
	http.HandleFunc("/", index)
//	http.HandleFunc("/health", health.Health)
//	mux.HandleFunc("/health", health.Health)

	http.ListenAndServe(":80", nil)
}
