package main

import (
	"cloud.google.com/go/spanner"
	"time"
	"fmt"
	"net/http"
	"log"
	"os"
	"context"
)

func main() {

	projectID, ok := os.LookupEnv("SPANNER_PROJECT_ID")
	if !ok {
		log.Fatalf("env not set: SPANNER_PROJECT_ID")
	}
	instanceID, ok := os.LookupEnv("SPANNER_INSTANCE_ID")
	if !ok {
		log.Fatalf("env not set: SPANNER_INSTANCE_ID")
	}
	databaseID, ok := os.LookupEnv("SPANNER_DATABASE_ID")
	if !ok {
		log.Fatalf("env not set: SPANNER_DATABASE_ID")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		dsn := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectID, instanceID, databaseID)
		client, err := spanner.NewClient(ctx, dsn)
		if err != nil {
			panic(err)
		}
		defer client.Close()

		start := time.Now()
		stmt := spanner.NewStatement("SELECT 1")
		client.Single().Query(ctx, stmt).Do(func(row *spanner.Row) error {
			return nil
		})
		fmt.Fprintf(w, "Query time: %s\n", time.Since(start))
	})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

