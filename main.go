package main

import (
	"log"

	"gats.dev/cloudsql-export-import/gcp"
)

func main() {
	log.Println("Hello, World!")

	client := gcp.GetCredentials()
	project_id := "bhd-platform-services-dev"
	databaseList := gcp.ListDatabases(project_id, client)

	for _, database := range databaseList.Items {
		log.Println(database.Name)
	}

	gcp.ListStorageBuckets()
	bucketHandle := gcp.GetStorageBucket("bhd-ccdata-dev")
	log.Println(bucketHandle.BucketName())
}
