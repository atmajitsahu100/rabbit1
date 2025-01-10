package main

import (
	"fmt"
	"net/http"
)

var secretKey1 = "0000000000012121"
var secretKey2 = "00002121"
var secretKey3 = "323242"

// The function always writes the greeting regardless of the request details.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// Returns true if the key matches any of the valid keys, otherwise returns false.
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

// authMiddleware is a middleware function that validates the Authorization header before allowing the request to proceed.
// It checks if a valid authorization key is present in the request header. If the key is missing or invalid,
// it returns a 401 Unauthorized error. Otherwise, it calls the next handler in the request processing chain.
//
// The function takes an http.Handler as input and returns a new http.Handler that wraps the original handler
// with authentication logic.
//
// Parameters:
//   - next: The next HTTP handler to be called if authentication is successful
//
// Returns:
//   - An http.Handler that performs authentication before calling the next handler
//
// Example:
//   http.Handle("/", authMiddleware(myHandler))
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
// It registers the root route ("/") with the handler wrapped by the authentication middleware.
// The server will validate incoming requests using the predefined secret keys before processing.
// If no errors occur during server startup, it will listen for and handle incoming HTTP requests.
// Note: This function will block and run indefinitely until the server is manually stopped.
func main() {
	http.Handle("/", authMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}
