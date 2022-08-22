package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/k1nho/goKingdomAPI/controllers/products_ctrl"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/kingdom/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/kingdom/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/kingdom/products/{product_id}", controllers.UpdateProduct).Methods("POST")
	router.HandleFunc("/api/kingdom/products/{product_id}", controllers.GetProduct).Methods("GET")
}

func RunServer() {
	router = mux.NewRouter()
	initHandlers()
	log.Fatal(http.ListenAndServe(":8000", router))
}
