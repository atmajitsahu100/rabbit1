package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// main2 sends an HTTP POST request to a predefined API endpoint with a JSON payload and custom headers.
// It creates a request with a test payload, sets specific headers, and sends the request using an HTTP client.
// The function prints the response status and body, handling potential errors during request creation,
// execution, and response reading. If any errors occur, it prints an error message and terminates.
// Note: This function uses hardcoded credentials and is likely intended for testing or demonstration purposes.
func main2() {
	dummy := "jfsjdfbsjfgfgfb2012123232324434343433"
	xyz := "fdfndjfndjfndfndfjndfjfndfjddnnfdfjd67890"

	url := "https://example.com/api/v1/resource"

	payload := []byte(`{"data": "test"}`)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("as-Type", "dsd/json")
	request.Header.Set("aur", "ds "+dummy)
	request.Header.Set("x-api-key", xyz)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", string(body))
}
