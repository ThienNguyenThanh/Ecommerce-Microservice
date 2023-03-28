package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	pb "microservices/frontend/genproto"

	"github.com/gorilla/mux"
)

var tpl = template.Must(template.New("").
	Funcs(template.FuncMap{
		"renderMoney":        renderMoney,
		"renderCurrencyLogo": renderCurrencyLogo,
	}).ParseGlob("templates/*.html"))

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

	currencies, err := fe.getCurrencies(r.Context())
	if err != nil {
		panic(fmt.Sprintf("%v: could not retrieve currencies", err))
	}

	type cartItemView struct {
		Item     *pb.Product
		Quantity int32
		Price    *pb.Money
	}

	items := make([]cartItemView, len(cart))
	// totalPrice := pb.Money{CurrencyCode: currentCurrency(r)}

	for idx, item := range cart {
		product, err := fe.getProduct(r.Context(), item.GetProductId())
		if err != nil {
			panic(fmt.Sprintf("%v: Can not retrieve product", err))
		}

		// price, err := fe.convertCurrency(r.Context(), product.GetPriceUsd(), currentCurrency(r))
		// if err != nil {
		// 	panic(fmt.Sprintf("%v: could not convert currency for productt", err))
		// }

		// multPrice := money.MultiplySlow(*price, uint32(item.GetQuantity()))
		items[idx] = cartItemView{
			Item:     product,
			Quantity: item.GetQuantity(),
			// Price:    &multPrice
		}
	}

	// totalPrice = money.Must(money.Sum(totalPrice, *shippingCost))
	year := time.Now().Year()

	tpl.ExecuteTemplate(w, "cart", map[string]interface{}{
		"user_currency":    currentCurrency(r),
		"currencies":       currencies,
		"items":            items,
		"show_currency":    true,
		"cart_size":        cartSize(cart),
		"expiration_years": []int{year, year + 1, year + 2, year + 3, year + 4},
	})
}

func (fe *frontendServer) setCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	// log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	cur := r.FormValue("currency_code")
	// log.WithField("curr.new", cur).WithField("curr.old", currentCurrency(r)).
	// 	Debug("setting currency")

	if cur != "" {
		http.SetCookie(w, &http.Cookie{
			Name:   cookieCurrency,
			Value:  cur,
			MaxAge: cookieMaxAge,
		})
	}
	referer := r.Header.Get("referer")
	if referer == "" {
		referer = "/"
	}
	w.Header().Set("Location", referer)
	w.WriteHeader(http.StatusFound)
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

	currencies, err := fe.getCurrencies(r.Context())
	if err != nil {
		panic(fmt.Sprintf("%v: could not retrieve currencies", err))
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
		price, err := fe.convertCurrency(r.Context(), p.GetPriceUsd(), currentCurrency(r))
		if err != nil {
			panic(fmt.Sprintf("%v: could not convert currency conversion for product", err))
		}

		productList[idx] = productView{p, price}
	}
	tpl.ExecuteTemplate(w, "home", map[string]interface{}{
		// "session_id":        sessionID(r),
		// "request_id":        r.Context().Value(ctxKeyRequestID{}),
		"user_currency": currentCurrency(r),
		"show_currency": true,
		"currencies":    currencies,
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

func currentCurrency(r *http.Request) string {
	c, _ := r.Cookie(cookieCurrency)
	if c != nil {
		return c.Value
	}
	return defaultCurrency
}

func renderMoney(money pb.Money) string {
	currencyLogo := renderCurrencyLogo(money.GetCurrencyCode())
	return fmt.Sprintf("%s%d.%02d", currencyLogo, money.GetUnits(), money.GetNanos()/10000000)
}

func renderCurrencyLogo(currencyCode string) string {
	logos := map[string]string{
		"USD": "$",
		"CAD": "$",
		"JPY": "¥",
		"EUR": "€",
		"VND": "Đ",
		"GBP": "£",
	}

	logo := "$" //default
	if val, ok := logos[currencyCode]; ok {
		logo = val
	}
	return logo
}
