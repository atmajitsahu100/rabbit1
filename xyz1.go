package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// main2 sends a POST request to a predefined API endpoint with a test payload and custom headers.
// It constructs an HTTP request with a JSON payload, sets specific headers including an API key,
// and sends the request to the specified URL. The function prints the response status and body,
// handling potential errors during request creation, execution, and response reading.
// If any errors occur during the process, the function logs the error and terminates.
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
