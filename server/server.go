// server/server.go
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// User はユーザーデータの構造体です。
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	// モックデータとして使用するユーザーのマップ
	// サーバーが起動している時だけデータが保持される
	users = map[string]User{
		"1": {ID: "1", Name: "Alice"},
		"2": {ID: "2", Name: "Bob"},
	}
	mu sync.Mutex
)

// handler はリクエストに応じてCRUD処理を行います。
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		createUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// createUser は新しいユーザーを作成します。
func createUser(w http.ResponseWriter, r *http.Request) {
	// lockについては要するにトランザクション処理を行なっている。
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	users[user.ID] = user
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
}

// getUser はユーザー情報を取得します。
func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	mu.Lock()
	user, exists := users[id]
	mu.Unlock()
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// updateUser は既存のユーザー情報を更新します。
func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	_, exists := users[user.ID]
	if !exists {
		mu.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	users[user.ID] = user
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
}

// deleteUser はユーザーを削除します。
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// クエリからidを取得
	id := r.URL.Query().Get("id")
	mu.Lock()
	_, exists := users[id]
	if !exists {
		mu.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(users, id)
	mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

// StartServer はHTTPサーバーを起動します。
func StartServer() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
