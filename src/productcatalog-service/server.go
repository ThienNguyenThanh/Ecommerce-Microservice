package main

import (
	// "bytes"
	"context"
	"fmt"
	pb "microservices/product/genproto"
	"net"
	"os"
	"strings"
	"time"
	// "encoding/gob"

	// "github.com/golang/protobuf/jsonpb"
	
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "google.golang.org/protobuf/encoding/protojson"
	// "github.com/redis/go-redis/v9"
)

var (
	cat pb.ListProductsResponse
	log *logrus.Logger

	port = "3050"
)

type productCatalog struct {
	pb.UnimplementedProductCatalogServiceServer
}

type PriceUsd struct {
	CurrencyCode string `bson:"currencyCode"`
	Units        int64
	Nanos        int32
}

type productCollection struct {
	Id          string `bson:"_id"`
	Name        string
	Description string
	Picture     string
	PriceUsd    PriceUsd `bson:"priceUsd"`
	Categories  []string
	InStock     int32
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

	// err := readProductFile(&cat)
	// if err != nil {
	// 	log.Warnf("could not parse product catalog")
	// }
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


func parseCatalog() []*pb.Product {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("Product-Service").Collection("Products")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	var results []productCollection
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	res := []*pb.Product{}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
		res = append(res, 
						&pb.Product{
							Id: result.Id,
							Name: result.Name,
							Description: result.Description,
							Picture: result.Picture,
							PriceUsd: &pb.Money{
								CurrencyCode: result.PriceUsd.CurrencyCode,
								Units: result.PriceUsd.Units,
								Nanos: result.PriceUsd.Nanos,
							},
							Categories: result.Categories,
							InStock: result.InStock,
						})
	}
	cat.Products = res
	return res
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
