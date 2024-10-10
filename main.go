package main

import (
	"go-transaction/model"
	"go-transaction/route"
)

func main() {

	db, _ := model.DBConnection()
	route.SetupRoutes(db)
}
