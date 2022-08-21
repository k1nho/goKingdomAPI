package main

import (
	"github.com/k1nho/goKingdomAPI/controllers"
	"github.com/k1nho/goKingdomAPI/model"
)

func main() {
	model.Init()
	controllers.RunServer()
}
