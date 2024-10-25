package gcp

import (
	"log"
	"net/http"

	"cloud.google.com/go/auth/credentials"
	"cloud.google.com/go/auth/httptransport"
)

func GetCredentials() *http.Client {
	// You can detect options here, not sure what they can be used for but still
	creds, err := credentials.DetectDefault(&credentials.DetectOptions{})
	if err != nil {
		log.Fatal(err)
	}
	client, err := httptransport.NewClient(&httptransport.Options{
		Credentials: creds,
	})
	if err != nil {
		log.Fatal(err)
	}
	return client
}
