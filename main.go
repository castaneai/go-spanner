package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc"
	channelzsvc "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	dsn, ok := os.LookupEnv("SPANNER_DSN")
	if !ok {
		log.Fatalf("env: SPANNER_DSN not set")
	}

	go func() {
		if err := setupChannelz(); err != nil {
			log.Fatal(err)
		}
	}()

	client, err := spanner.NewClient(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			start := time.Now()
			stmt := spanner.NewStatement("SELECT 1")
			if err := client.Single().Query(ctx, stmt).Do(func(row *spanner.Row) error {
				return nil
			}); err != nil {
				log.Fatal(err)
			}
			log.Printf("Query time: %s\n", time.Since(start))
		}
	}
}

func setupChannelz() error {
	s := grpc.NewServer()
	reflection.Register(s)
	channelzsvc.RegisterChannelzServiceToServer(s)
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		return err
	}
	log.Printf("Channelz server listening on :8000...")
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
