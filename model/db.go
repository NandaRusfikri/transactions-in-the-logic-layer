package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConnection -> return db instance
func DBConnection() (*gorm.DB, error) {
	USER := "admin"
	PASS := "admin"
	HOST := "localhost"
	PORT := "6543"
	DBNAME := "db-transaction"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", HOST, USER, PASS, DBNAME, PORT)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

}
