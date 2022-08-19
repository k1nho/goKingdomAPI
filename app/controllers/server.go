package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func RunServer() {
	router = mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
