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

	// runListProducts(client)
	// runGetProduct(client, "OLJCESPC7Z")
	runSearchProducts(client, "mug")
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

func runGetProduct(client pb.ProductCatalogServiceClient, productId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	product, err := client.GetProduct(ctx, &pb.GetProductRequest{Id: productId})
	if err != nil {
		log.Fatalf("Did not find any product: %v", err)
	}
	log.Printf("Found: %v", product)
}

func runSearchProducts(client pb.ProductCatalogServiceClient, query string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	found, err := client.SearchProducts(ctx, &pb.SearchProductsRequest{Query: query})
	if err != nil {
		log.Fatalf("Did not find any product: %v", err)
	}
	log.Printf("Result: %v", found)
}
