package main

import (
	"log"

	"gats.dev/cloudsql-export-import/gcp"
)

func main() {
	log.Println("Creating a call that starts an export")

	client := gcp.GetCredentials()
	project_id := "bhd-platform-services-dev"
	// databaseList := gcp.ListDatabases(project_id, client)

	// for _, database := range databaseList.Items {
	// 	log.Println(database.Name)
	// }

	// gcp.ListStorageBuckets()
	// bucketHandle := gcp.GetStorageBucket("bhd-ccdata-dev")
	// log.Println(bucketHandle.BucketName())
	exportBucketName := "vision-db-instance-dev-backup"
	exportDBInstanceName := "vision-db-instance-development"
	exportDBName := "test_cm"
	gcp.StartExport(
		project_id,
		client,
		exportBucketName,
		exportDBName,
		exportDBInstanceName,
	)
}
