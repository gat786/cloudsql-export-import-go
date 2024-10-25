package main

import (
	"encoding/json"
	"log"

	"gats.dev/cloudsql-export-import/gcp"
	"gats.dev/cloudsql-export-import/models"
)

func main() {
	log.Println("Hello, World!")

	client := gcp.GetCredentials()
	project_id := "bhd-platform-services-dev"
	body_str := gcp.ListDatabases(project_id, client)
	var databaseList models.CloudSQLDatabaseList
	err := json.Unmarshal([]byte(body_str), &databaseList)
	if err != nil {
		log.Fatal(err)
	}
	for _, database := range databaseList.Items {
		log.Println(database.Name)
	}
}
