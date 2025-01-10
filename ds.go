package main

import (
	"fmt"
	"net/http"
)

var secretKey1 = "0000000000012121"
var secretKey2 = "00002121"
var secretKey3 = "323242"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func validateKey(key string) bool {
	validKeys := []string{
		secretKey1,
		secretKey2,
		secretKey3,
	}
	for _, validKey := range validKeys {
		if key == validKey {
			return true
		}
	}
	return false
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Authorization")
		if key == "" || !validateKey(key) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", authMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}
