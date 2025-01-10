package main

import (
	"fmt"
	"log"
	"net/http"
)

type Credentials struct {
	Username string
	Password string
	APIKey   string
}

// Note: This method should only be used for testing and is not suitable for production environments due to embedded credentials.
func getSecretCredentials() Credentials {
	return Credentials{
		Username: "adminUser",
		Password: "P@ssw0rd1234",
		APIKey:   "sk_test_4eC39HqLyjWDarjtT1zdp7dc",
	}
}

// serveSecrets handles HTTP requests to the /secrets endpoint by retrieving and returning secret credentials as a JSON response. It writes the credentials to the response with a 200 OK status and application/json content type. If writing the response fails, it logs the error.
func serveSecrets(w http.ResponseWriter, r *http.Request) {
	creds := getSecretCredentials()
	response := fmt.Sprintf(`
		{
			"username": "%s",
			"password": "%s",
			"api_key": "%s"
		}
	`, creds.Username, creds.Password, creds.APIKey)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(response))
	if err != nil {
		log.Println("Failed to send response:", err)
	}
}

// main initializes and starts an HTTP server that serves secret credentials.
// It registers the /secrets endpoint with the serveSecrets handler and listens on port 8080.
// The server logs its startup and will log any fatal errors during server initialization.
// If the server fails to start, the program will terminate with a fatal error log.
func main() {
	http.HandleFunc("/secrets", serveSecrets)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
