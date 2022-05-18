package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/lorenzo-milicia/go-client-queue/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewDataFetcherClient(conn)

	stream, err := client.FetchQueueStream(context.Background(), &api.StreamSize{Size: 1000})

	if err != nil {
		log.Fatal(err)
	}

	for {
		batch, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received stream of %v records\n", len(batch.Records))
	}
}
