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

// getSecretCredentials returns a Credentials struct with hardcoded secret values for username, password, and API key. This function is primarily used for demonstration or testing purposes and should not be used in production environments due to the exposure of sensitive credentials.
func getSecretCredentials() Credentials {
	return Credentials{
		Username: "adminUser",
		Password: "P@ssw0rd1234",
		APIKey:   "sk_test_4eC39HqLyjWDarjtT1zdp7dc",
	}
}

// serveSecrets handles HTTP requests to the /secrets endpoint by retrieving and returning secret credentials as a JSON response. It writes the credentials to the response writer and logs any errors encountered during the response writing process.
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

// The server will log a startup message and will terminate with a fatal error if the server fails to start.
func main() {
	http.HandleFunc("/secrets", serveSecrets)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
