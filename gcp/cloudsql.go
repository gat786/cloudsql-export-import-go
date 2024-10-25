package gcp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gats.dev/cloudsql-export-import/models"
)

var (
	baseurl = "https://sqladmin.googleapis.com"
)

func ListDatabases(project_name string, client *http.Client) models.CloudSQLDatabaseListResponse {
	url := fmt.Sprintf("%s/v1/projects/%s/instances", baseurl, project_name)
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
	var databaseList models.CloudSQLDatabaseListResponse
	decodingErr := json.Unmarshal([]byte(body_str), &databaseList)
	if decodingErr != nil {
		log.Fatal(decodingErr)
	}
	return databaseList
}
