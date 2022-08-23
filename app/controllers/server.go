package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	productsRouter := router.PathPrefix("/api/kingdom/products").Subrouter()
	productsRouter.HandleFunc("", GetProducts).Methods("GET")
	productsRouter.HandleFunc("", CreateProduct).Methods("POST")
	productsRouter.HandleFunc("/{product_id}", UpdateProduct).Methods("PUT")
	productsRouter.HandleFunc("/{product_id}", GetProduct).Methods("GET")
	productsRouter.HandleFunc("/{product_id}", DeleteProduct).Methods("DELETE")
	productsRouter.HandleFunc("/batch", FetchProductBatch).Methods("POST")
}

func RunServer() {
	router = mux.NewRouter()
	initHandlers()
	log.Fatal(http.ListenAndServe(":8000", router))
}
