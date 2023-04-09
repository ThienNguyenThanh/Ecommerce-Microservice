package main

import (
	"context"
	"fmt"
	pb "microservices/order/genproto"
	"net"
	"os"
	"time"

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
	pb.UnimplementedOrderServiceServer
}

type Address struct {
	StreetAddress string `bson:"street_address"`
	City          string
	State         string `bson:"omitempty"`
	Country       string
	ZipCode       int32 `bson:"zip_code"`
}
type Money struct {
	CurrencyCode string `bson:"currency_code"`
	Units        int64
	Nanos        int32
}
type CartItem struct {
	ProductId string `bson:"product_id"`
	Quantity  int32
}
type OrderItem struct {
	Item CartItem
	Cost Money
}

type OrderCollection struct {
	OrderId            string      `bson:"order_id"`
	ShippingTrackingId string      `bson:"shipping_tracking_id"`
	ShippingCost       Money       `bson:"shipping_cost"`
	ShippingAddress    Address     `bson:"shipping_address"`
	Items              []OrderItem `bson:"items"`
	DeliveryStatus     string      `bson:"delivery_status"`
	PaymentStatus      string      `bson:"payment_status"`
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
	// 	log.Warnf("could not insert order.")
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

	svc := &Order{}

	pb.RegisterOrderServiceServer(srv, svc)
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
	newOrder := OrderCollection{
		OrderId:            "order1",
		ShippingTrackingId: "tracking-ID-1",
		ShippingCost: Money{
			CurrencyCode: "USD",
			Units:        19,
			Nanos:        990000000,
		},
		ShippingAddress: Address{
			StreetAddress: "100 Dien Bien Phu",
			City:          "Ho Chi Minh",
			State:         "N/A",
			Country:       "VN",
			ZipCode:       10000,
		},
		Items: []OrderItem{
			OrderItem{
				Item: CartItem{
					ProductId: "ipad_pro",
					Quantity:  1,
				},
				Cost: Money{
					CurrencyCode: "USD",
					Units:        19,
					Nanos:        990000000,
				},
			},
			OrderItem{
				Item: CartItem{
					ProductId: "ipad_pro",
					Quantity:  1,
				},
				Cost: Money{
					CurrencyCode: "USD",
					Units:        19,
					Nanos:        990000000,
				},
			},
		},
		DeliveryStatus: "Pending",
		PaymentStatus:  "paid",
	}

	result, err := collection.InsertOne(context.TODO(), newOrder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}
