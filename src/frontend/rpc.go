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

func (fe *frontendServer) getCart(ctx context.Context, userID string) ([]*pb.CartItem, error) {
	response, err := pb.NewCartServiceClient(fe.cartServiceConn).
		GetCart(ctx, &pb.GetCartRequest{UserId: userID})

	return response.GetItems(), err
}

// func (fe *frontendServer) getCurrencies(ctx context.Context) ([]string, error) {
// 	currs, err := pb.NewCurrencyServiceClient(fe.currencySvcConn).
// 		GetSupportedCurrencies(ctx, &pb.Empty{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	var out []string
// 	for _, c := range currs.CurrencyCodes {
// 		if _, ok := whitelistedCurrencies[c]; ok {
// 			out = append(out, c)
// 		}
// 	}
// 	return out, nil
// }
