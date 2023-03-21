package main

import (
	"bytes"
	"context"
	"fmt"
	pb "microservices/product/genproto"
	"net"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	cat pb.ListProductsResponse
	log *logrus.Logger

	port = "3050"
)

type productCatalog struct {
	pb.UnimplementedProductCatalogServiceServer
}

func init() {
	log = logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout
	err := readProductFile(&cat)
	if err != nil {
		log.Warnf("could not parse product catalog")
	}
}

func main() {
	log.Printf("Start server at: %v", port)
	run(port)
}

func run(port string) string {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	var srv *grpc.Server = grpc.NewServer()

	svc := &productCatalog{}

	pb.RegisterProductCatalogServiceServer(srv, svc)
	srv.Serve(l)

	return l.Addr().String()
}

func readProductFile(products *pb.ListProductsResponse) error {
	productJSON, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatalf("failed to open product catalog json file: %v", err)
		return err
	}
	if err := jsonpb.Unmarshal(bytes.NewReader(productJSON), products); err != nil {
		log.Warnf("failed to parse the catalog JSON: %v", err)
		return err
	}
	log.Info("successfully parsed product catalog json")
	return nil
}

func parseCatalog() []*pb.Product {
	if len(cat.Products) == 0 {
		err := readProductFile(&cat)
		if err != nil {
			return []*pb.Product{}
		}
	}
	return cat.Products
}

func (p *productCatalog) ListProducts(context.Context, *pb.Empty) (
	*pb.ListProductsResponse, error) {
	return &pb.ListProductsResponse{Products: parseCatalog()}, nil
}

func (p *productCatalog) GetProduct(ctx context.Context, req *pb.GetProductRequest) (
	*pb.Product, error) {
	var found *pb.Product
	for i := 0; i < len(parseCatalog()); i++ {
		if req.Id == parseCatalog()[i].Id {
			found = parseCatalog()[i]
		}
	}
	if found == nil {
		return nil, status.Errorf(codes.NotFound, "no product with ID %s", req.Id)
	}
	return found, nil
}

func (p *productCatalog) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (
	*pb.SearchProductsResponse, error) {
	var ps []*pb.Product
	for _, p := range parseCatalog() {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(req.Query)) ||
			strings.Contains(strings.ToLower(p.Description), strings.ToLower(req.Query)) {
			ps = append(ps, p)
		}
	}
	return &pb.SearchProductsResponse{Results: ps}, nil
}
