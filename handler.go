package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func (fe *frontendServer) headerHandler(w http.ResponseWriter, r *http.Request) {
	// p := Header{Title: "Header", User: "Thien"}
	tpl.ExecuteTemplate(w, "header", nil)
}

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())
	tpl.ExecuteTemplate(w, "home", map[string]interface{}{
		// "session_id":        sessionID(r),
		// "request_id":        r.Context().Value(ctxKeyRequestID{}),
		// "user_currency":     currentCurrency(r),
		"show_currency": true,
		"currencies":    "VND",
		// "products":          ps,
		// "cart_size":         cartSize(cart),
		// "banner_color":      os.Getenv("BANNER_COLOR"), // illustrates canary deployments
		// "ad":                fe.chooseAd(r.Context(), []string{}, log),
		// "platform_css":      plat.css,
		// "platform_name":     plat.provider,
		// "is_cymbal_brand":   isCymbalBrand,
		// "deploymentDetails": deploymentDetailsMap,
	})
}
