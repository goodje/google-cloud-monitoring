package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
)

func main() {
	obtainOauth2Token()
}

func obtainOauth2Token() {
	ctx := context.Background()

	//  scopes := []string{
	//  	"https://www.googleapis.com/auth/cloud-platform",
	//  }

	// Use Application Default Credentials (ADC)
	creds, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		log.Printf("Unable to find default credentials: %v", err)

		jsonKeyPath := "path/to/your/service-account-key.json"

		// Read the service account JSON key file
		data, err := ioutil.ReadFile(jsonKeyPath)
		if err != nil {
			log.Fatalf("Unable to read the service account key file: %v", err)
		}

		// Obtain a token source based on the service account JSON key
		creds, err = google.CredentialsFromJSON(ctx, data)
		if err != nil {
			log.Fatalf("Unable to obtain credentials from JSON key: %v", err)
		}
	}

	// Get a token from the token source
	token, err := creds.TokenSource.Token()
	if err != nil {
		log.Fatalf("Unable to get token: %v", err)
	}

	// Print the obtained token
	fmt.Printf("Access Token: %s\n", token.AccessToken)
	fmt.Printf("Token Expiry: %s\n", token.Expiry)
}
