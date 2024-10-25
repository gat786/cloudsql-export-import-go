package gcp

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func ListDatabases(project_name string, client *http.Client) string {
	url := fmt.Sprintf("https://sqladmin.googleapis.com/v1/projects/%s/instances", project_name)
	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// this will close the response body when the function returns
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	body_str := string(body)
	return body_str
}
