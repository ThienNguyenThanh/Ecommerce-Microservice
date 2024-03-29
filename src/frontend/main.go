package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"github.com/sirupsen/logrus"
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
		"VND": true}
)

const (
	port            = "8080"
	defaultCurrency = "USD"
	cookieMaxAge    = 60 * 60 * 48

	cookiePrefix    = "shop_"
	cookieSessionID = cookiePrefix + "session-id"
	cookieCurrency  = cookiePrefix + "currency"
)

type ctxKeySessionID struct {}

type frontendServer struct {
	productCatalogServiceAddress string
	productCatalogServiceConn    *grpc.ClientConn

	cartServiceAddress string
	cartServiceConn    *grpc.ClientConn

	currencyServiceAddress string
	currencyServiceConn    *grpc.ClientConn

	shippingServiceAddress string
	shippingServiceConn    *grpc.ClientConn
}

func main() {

	ctx := context.Background()
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

	srvPort := port
	addr := os.Getenv("LISTEN_ADDR")
	svc := new(frontendServer)

	mustMapEnv(&svc.productCatalogServiceAddress, "PRODUCT_CATALOG_SERVICE_ADDR")
	mustMapEnv(&svc.cartServiceAddress, "CART_SERVICE_ADDR")
	mustMapEnv(&svc.currencyServiceAddress, "CURRENCY_SERVICE_ADDR")
	mustMapEnv(&svc.shippingServiceAddress, "SHIPPING_SERVICE_ADDR")

	mustConnGRPC(ctx, &svc.productCatalogServiceConn, svc.productCatalogServiceAddress)
	mustConnGRPC(ctx, &svc.cartServiceConn, svc.cartServiceAddress)
	mustConnGRPC(ctx, &svc.currencyServiceConn, svc.currencyServiceAddress)
	mustConnGRPC(ctx, &svc.shippingServiceConn, svc.shippingServiceAddress)

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/", svc.homeHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/product/{id}", svc.productHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/cart", svc.viewCartHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/cart", svc.addToCartHandler).Methods(http.MethodPost)
	r.HandleFunc("/cart/empty", svc.emptyCartHandler).Methods(http.MethodPost)
	r.HandleFunc("/setCurrency", svc.setCurrencyHandler).Methods(http.MethodPost)

	var handler http.Handler = r
	handler = &logHandler{log: log, next: handler}     // add logging
	handler = ensureSessionID(handler)                 // add session ID

	log.Infof("Start server on " + addr + ":" + srvPort)
	log.Fatal(http.ListenAndServe(addr+":"+srvPort, handler))

	// viperenv := mustMapEnv("STRONGEST_AVENGER")

	// fmt.Printf("viper : %s = %s \n", "STRONGEST_AVENGER", viperenv)
}

func mustMapEnv(target *string, envKey string) {
	value := os.Getenv(envKey)
	if value == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}

	*target = value
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
