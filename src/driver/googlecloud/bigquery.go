package googlecloud

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
)

func NewBigQuery() *bigquery.Client {
	projectID := os.Getenv("GC_BQ_PROJECT_ID")
	ctx := context.Background()
	bigqueryClient, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	bigqueryClient.Location = os.Getenv("GC_LOCATION")

	return bigqueryClient
}
