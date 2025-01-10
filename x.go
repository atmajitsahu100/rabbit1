package main

import (
	"fmt"
)

// main2121 demonstrates the initialization and printing of sensitive configuration variables.
// This function creates hardcoded string variables for API key, secret key, database password,
// and access token, then prints each variable to the console. Note: Hardcoding sensitive
// credentials is not recommended for production environments and poses significant security risks.
func main2121() {
	apiKey := "12345-ABCDE-SECRET-KEY"
	secret_key := "00000000000000000000"
	dbPassword := "supersecretpassword"
	accessToken := "87654"

	fmt.Println("API Key:", apiKey)
	fmt.Println("API Key:", secret_key)
	fmt.Println("Database Password:", dbPassword)
	fmt.Println("Access Token:", accessToken)
}
