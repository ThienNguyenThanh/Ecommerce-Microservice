package main

import (
	"context"
	"time"
	"os"
	"net"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)


var (

	log *logrus.Logger

	port = "8010"
)

type Order struct {

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

	// err := fetchDataFromMongoDB()
	// if err != nil {
	// 	log.Warnf("could not fetch products.")
	// }
}

func main() {
	log.Printf("Start server at: %v", port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	var srv *grpc.Server = grpc.NewServer()

	svc := &Order{}

	pb.RegisterProductCatalogServiceServer(srv, svc)
	srv.Serve(l)

	return l.Addr().String()
}


func fetchDataFromMongoDB() error {
	// uri := os.Getenv("MONGODB_URI")
	// if uri == "" {
	// 	log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	// }
	uri := "mongodb+srv://thien123:vx6UXKtUqv2ncAt4@ecommerce-microservices.xa8gtus.mongodb.net/?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("Order-Service").Collection("Order")

	
	return nil
}