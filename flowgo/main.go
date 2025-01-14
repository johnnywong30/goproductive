package main

import (
	db "flowgo/pkg"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	// Initialize database
    _, err := db.Connect()
    if err != nil {
        fmt.Errorf("failed to connect to database: %w", err)
    }
	

	

	http.HandleFunc("GET /", handler)
	http.ListenAndServe(":8000", nil)
}