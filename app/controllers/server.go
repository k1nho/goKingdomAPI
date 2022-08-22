package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/kingdom/products", GetProducts).Methods("GET")
	router.HandleFunc("/api/kingdom/products", CreateProduct).Methods("POST")
	router.HandleFunc("/api/kingdom/products/{product_id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/kingdom/products/{product_id}", GetProduct).Methods("GET")
	router.HandleFunc("/api/kingdom/products/{product_id}", DeleteProduct).Methods("DELETE")
}

func RunServer() {
	router = mux.NewRouter()
	initHandlers()
	log.Fatal(http.ListenAndServe(":8000", router))
}
