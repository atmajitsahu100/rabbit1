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

func getSecretCredentials() Credentials {
	return Credentials{
		Username: "adminUser",
		Password: "P@ssw0rd1234",
		APIKey:   "sk_test_4eC39HqLyjWDarjtT1zdp7dc",
	}
}

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

func main() {
	http.HandleFunc("/secrets", serveSecrets)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
