package kv

import (
	"fmt"
	"net/http"
)

// Main is the main function for the kvStore
func Main() {
	fmt.Println("kvStore server runnning...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received request\n")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	// fmt.Fprintf(w, "Method: %s\n", r)
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		fmt.Fprintf(w, "GET")
	case http.MethodPost:
		fmt.Fprintf(w, "POST")
	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE")
	default:
		fmt.Fprintf(w, "default")
	}
}
