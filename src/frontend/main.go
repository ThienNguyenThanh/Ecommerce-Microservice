package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Header struct {
	Title string
	User  string
}

var (
	whitelistedCurrencies = map[string]bool{
		"USD": true,
		"EUR": true,
		"CAD": true,
		"JPY": true,
		"GBP": true,
		"TRY": true}
)

type frontendServer struct {
	productCatalogServiceAddress string
	productCatalogServiceConn    *grpc.ClientConn

	cartServiceAddress string
	cartServiceConn    *grpc.ClientConn
}

func main() {

	ctx := context.Background()
	srvPort := "3080"
	svc := new(frontendServer)

	mustMapEnv(&svc.productCatalogServiceAddress, "PRODUCT_CATALOG_SERVICE_ADDR")
	mustMapEnv(&svc.cartServiceAddress, "CART_SERVICE_ADDR")

	mustConnGRPC(ctx, &svc.productCatalogServiceConn, svc.productCatalogServiceAddress)
	mustConnGRPC(ctx, &svc.cartServiceConn, svc.cartServiceAddress)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", svc.homeHandler)
	http.HandleFunc("/header", svc.headerHandler)
	http.HandleFunc("/product", svc.productHandler)
	http.HandleFunc("/cart", svc.viewCartHandler)

	fmt.Println("Start server at port " + srvPort)
	log.Fatal(http.ListenAndServe("localhost:"+srvPort, nil))

	// viperenv := mustMapEnv("STRONGEST_AVENGER")

	// fmt.Printf("viper : %s = %s \n", "STRONGEST_AVENGER", viperenv)
}

func mustMapEnv(target *string, envKey string) {
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(envKey).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	// return value
	*target = value

	//   v := os.Getenv(envKey)
	// 	if v == "" {
	// 		panic(fmt.Sprintf("environment variable %q not set", envKey))
	// 	}
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("%v grpc: failed to connect %s", err, addr))
	}
}
