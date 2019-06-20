package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/spanner"
)

func main() {
	ctx := context.Background()
	dsn, ok := os.LookupEnv("SPANNER_DSN")
	if !ok {
		log.Fatalf("env: SPANNER_DSN not set")
	}
	client, err := spanner.NewClient(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	start := time.Now()
	stmt := spanner.NewStatement("SELECT 1")
	if err := client.Single().Query(ctx, stmt).Do(func(row *spanner.Row) error {
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	log.Printf("Query time: %s\n", time.Since(start))
}
