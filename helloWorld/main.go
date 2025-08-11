package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	auth := flag.String("auth", "", "authorization header value")
	flag.Parse()

	// Create request with authorization header
	req, err := http.NewRequest("POST", "https://auth.meshbot.co/oauth/token?grant_type=client_credentials", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	if *auth != "" {
		req.Header.Set("Authorization", *auth)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("%s", string(body))
}
