package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	pb "microservices/frontend/genproto"

	"github.com/gorilla/mux"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func (fe *frontendServer) addToCartHandler(w http.ResponseWriter, r *http.Request) {
	quantity, _ := strconv.ParseUint(r.FormValue("quantity"), 10, 32)
	productID := r.FormValue("product_id")
	if productID == "" || quantity == 0 {
		fmt.Println("Invalid form input")
	}

	product, err := fe.getProduct(r.Context(), productID)
	if err != nil {
		panic(fmt.Sprintf("%v: Can not retrieve product", err))
	}

	if err := fe.addToCart(r.Context(), "thien123", product.GetId(), int32(quantity)); err != nil {
		panic(fmt.Sprintf("%v: Fail to add to cart", err))
	}

	w.Header().Set("location", "/cart")
	w.WriteHeader(http.StatusFound)
}

func (fe *frontendServer) emptyCartHandler(w http.ResponseWriter, r *http.Request) {
	if err := fe.emptyCart(r.Context(), "thien123"); err != nil {
		panic(fmt.Sprintf("%v: Fail to add to cart", err))
	}

	w.Header().Set("location", "/cart")
	w.WriteHeader(http.StatusFound)
}

func (fe *frontendServer) viewCartHandler(w http.ResponseWriter, r *http.Request) {
	cart, err := fe.getCart(r.Context(), "thien123")
	if err != nil {
		panic(fmt.Sprintf("%v: could not retrieve cart", err))
	}

	type cartItemView struct {
		Item     *pb.Product
		Quantity int32
	}

	items := make([]cartItemView, len(cart))
	for idx, item := range cart {
		product, err := fe.getProduct(r.Context(), item.GetProductId())
		if err != nil {
			panic(fmt.Sprintf("%v: Can not retrieve product", err))
		}
		items[idx] = cartItemView{
			Item:     product,
			Quantity: item.GetQuantity(),
		}
	}

	tpl.ExecuteTemplate(w, "cart", map[string]interface{}{
		"items":     items,
		"cart_size": cartSize(cart),
	})
}

func (fe *frontendServer) productHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		panic(fmt.Sprintf("Can not load product %v", id))
	}

	productRetrieved, err := fe.getProduct(r.Context(), id)
	if err != nil {
		panic(fmt.Sprintf("%v: Can not retrieve product", err))
	}
	cart, err := fe.getCart(r.Context(), "thien123")
	if err != nil {
		panic(fmt.Sprintf("%v: could not retrieve cart", err))
	}

	product := struct {
		Item  *pb.Product
		Price *pb.Money
	}{productRetrieved, productRetrieved.GetPriceUsd()}
	tpl.ExecuteTemplate(w, "product", map[string]interface{}{
		"product":   product,
		"cart_size": cartSize(cart),
	})
}

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())

	products, err := fe.getProducts(r.Context())
	if err != nil {
		panic(fmt.Sprintf("%v: could not retrieve products", err))
	}

	cart, err := fe.getCart(r.Context(), "thien123")
	if err != nil {
		panic(fmt.Sprintf("%v: could not retrieve cart", err))
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
		"cart_size":     cartSize(cart),
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

// get total # of items in cart
func cartSize(c []*pb.CartItem) int {
	cartSize := 0
	for _, item := range c {
		cartSize += int(item.GetQuantity())
	}
	return cartSize
}
