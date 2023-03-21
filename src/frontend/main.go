package main

import (
	"fmt"
	"log"
	"net/http"
)

type Header struct {
	Title string
	User  string
}

type frontendServer struct {
}

func main() {

	svc := new(frontendServer)

	http.HandleFunc("/", svc.homeHandler)
	http.HandleFunc("/header", svc.headerHandler)

	fmt.Println("Start server at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
