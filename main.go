package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SecretHandler struct {
	SecretKey string
}

func NewSecretHandler(key string) *SecretHandler {
	return &SecretHandler{SecretKey: key}
}

func (h *SecretHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/get-secret" {
		h.handleSecretRequest(w, r)
		return
	}
	http.NotFound(w, r)
}

func (h *SecretHandler) handleSecretRequest(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"secret": h.SecretKey}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main1() {
	secretKey := "adnsjfsdfffdsssdsdcxcsdsuuudsfsfdsfdfdfsaddada12345"

	mux := http.NewServeMux()
	secretHandler := NewSecretHandler(secretKey)
	mux.Handle("/get-secret", secretHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
