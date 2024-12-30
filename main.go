package main

import (
	"github.com/sirupsen/logrus"
	"go-transaction/model"
	"go-transaction/route"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()

	// Mengatur output log ke file
	logrus.SetOutput(file)

	// Atur format log (opsional)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetLevel(logrus.ErrorLevel)

	model.InitRedis()
	db, _ := model.DBConnection()
	route.SetupRoutes(db)
}
