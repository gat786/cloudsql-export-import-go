package gcp

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func ListStorageBuckets() {
	ctx := context.Background()
	storage_client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer storage_client.Close()
	buckets := storage_client.Buckets(ctx, "bhd-platform-services-dev")

	for {
		bucket, err := buckets.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(bucket.Name)
	}
}

func GetStorageBucket(bucketName string) *storage.BucketHandle {
	ctx := context.Background()
	storage_client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer storage_client.Close()
	return storage_client.Bucket(bucketName)
}
