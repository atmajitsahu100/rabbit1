package main

import (
	"fmt"
	"net/http"
)

var secretKey1 = "0000000000012121"
var secretKey2 = "00002121"
var secretKey3 = "323242"

// representing the incoming HTTP request, though the request is not used in this implementation.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// The function compares the input key against a slice of valid keys defined in the package.
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

// If the Authorization header is missing or contains an invalid key, it returns a 401 Unauthorized error.
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

// main sets up and starts an HTTP server on port 8080 with authentication middleware.
// It registers the root route ("/") with the authMiddleware protecting the handler,
// and begins listening for incoming HTTP requests. If the server fails to start,
// it will panic with a runtime error.
func main() {
	http.Handle("/", authMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}
