package gcp

import (
	"bytes"
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

func StartExport(project_name string, client *http.Client, bucket_name string, database_name string, instance_name string) {
	url := fmt.Sprintf("%s/v1/projects/%s/instances/%s/export", baseurl, project_name, instance_name)
	exportContext := models.ExportContext{
		FileType: "SQL",
		URI:      fmt.Sprintf("gs://%s/%s.sql", bucket_name, database_name),
		Databases: []string{
			database_name,
		},
		Offload: false,
		SqlExportOptions: struct {
			Clean    bool `json:"clean"`
			IfExists bool `json:"ifExists"`
		}{
			Clean:    true,
			IfExists: true,
		},
	}

	exportRequestBody := models.ExportRequestBody{
		ExportContext: exportContext,
	}

	jsonBytes, err := json.Marshal(exportRequestBody)
	if err != nil {
		log.Fatal(err)
	}

	jsonBody := string(jsonBytes)
	log.Println(jsonBody)

	req, err := http.NewRequest("POST", url, nil)
	// set payload as body of the request
	req.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	body_str := string(body)
	log.Println(body_str)
}
