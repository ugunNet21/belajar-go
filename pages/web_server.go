// src/web_server.go

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a simple web server!")
}

func main() {
	// Menetapkan route ke handler
	http.HandleFunc("/", handler)

	// Menjalankan server pada port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
