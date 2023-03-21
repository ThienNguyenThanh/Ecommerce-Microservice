package main

import (
	"context"
	"log"
	pb "microservices/product/genproto"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:3050"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewProductCatalogServiceClient(conn)

	runListProducts(client)
}

func runListProducts(client pb.ProductCatalogServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ListProducts(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Did not find any product: %v", err)
	}

	log.Printf("List of products: %v", res)
}
