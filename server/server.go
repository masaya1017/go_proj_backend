// server/server.go
package server

import (
	"fmt"
	"net/http"
)

// StartServer はHTTPサーバーを開始する関数です。
func StartServer() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// handler はHTTPリクエストに応答する関数です。
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go web server!")
}
