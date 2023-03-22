package main

import (
	"context"
	pb "microservices/frontend/genproto"
)

func (fe *frontendServer) getProducts(ctx context.Context) ([]*pb.Product, error) {
	products, err := pb.NewProductCatalogServiceClient(fe.productCatalogServiceConn).
		ListProducts(ctx, &pb.Empty{})

	return products.GetProducts(), err
}

func (fe *frontendServer) getProduct(ctx context.Context, id string) (*pb.Product, error) {
	product, err := pb.NewProductCatalogServiceClient(fe.productCatalogServiceConn).
		GetProduct(ctx, &pb.GetProductRequest{Id: id})

	return product, err
}
