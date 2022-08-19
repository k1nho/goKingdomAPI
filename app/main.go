package main

import (
	"log"

	"github.com/k1nho/goKingdomAPI/controllers"
	"github.com/k1nho/goKingdomAPI/model"
)

func main() {
	err := model.Init()
	if err != nil {
		log.Fatal("could not initialize db")
		return
	}

	controllers.RunServer()
}
