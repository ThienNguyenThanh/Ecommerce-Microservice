package main

import (
	"fmt"
	"html/template"
	"net/http"

	pb "microservices/frontend/genproto"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func (fe *frontendServer) headerHandler(w http.ResponseWriter, r *http.Request) {
	// p := Header{Title: "Header", User: "Thien"}
	tpl.ExecuteTemplate(w, "header", nil)
}

func (fe *frontendServer) productHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		panic(fmt.Sprintf("Can not load product %v", id))
	}

	productRetrieved, err := fe.getProduct(r.Context(), id)
	if err != nil {
		panic(fmt.Sprintf("%v: Can not retrieve product", err))
	}

	product := struct {
		Item  *pb.Product
		Price *pb.Money
	}{productRetrieved, productRetrieved.GetPriceUsd()}
	tpl.ExecuteTemplate(w, "product", map[string]interface{}{
		"product": product,
	})
}

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())

	products, err := fe.getProducts(r.Context())
	if err != nil {
		panic(fmt.Sprintf("%v: Can not load product", err))
	}

	type productView struct {
		Item  *pb.Product
		Price *pb.Money
	}
	productList := make([]productView, len(products))
	for idx, p := range products {
		productList[idx] = productView{p, p.GetPriceUsd()}
	}
	tpl.ExecuteTemplate(w, "home", map[string]interface{}{
		// "session_id":        sessionID(r),
		// "request_id":        r.Context().Value(ctxKeyRequestID{}),
		// "user_currency":     currentCurrency(r),
		"show_currency": true,
		"currencies":    "VND",
		"products":      productList,
		// "cart_size":         cartSize(cart),
		// "banner_color":      os.Getenv("BANNER_COLOR"), // illustrates canary deployments
		// "ad":                fe.chooseAd(r.Context(), []string{}, log),
		// "platform_css":      plat.css,
		// "platform_name":     plat.provider,
		// "is_cymbal_brand":   isCymbalBrand,
		// "deploymentDetails": deploymentDetailsMap,
	})
}

// func renderMoney(money *pb.Money) string {
// 	currencyLogo := renderCurrencyLogo(money.GetCurrencyCode())
// 	return fmt.Sprintf("%s%d.%02d", currencyLogo, money.GetUnits(), money.GetNanos()/10000000)
// }

// func renderCurrencyLogo(currencyCode string) string {
// 	logos := map[string]string{
// 		"USD": "$",
// 		"CAD": "$",
// 		"JPY": "¥",
// 		"EUR": "€",
// 		"TRY": "₺",
// 		"GBP": "£",
// 	}

// 	logo := "$" //default
// 	if val, ok := logos[currencyCode]; ok {
// 		logo = val
// 	}
// 	return logo
// }
