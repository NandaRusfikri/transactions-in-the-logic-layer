package repository

import (
	"go-transaction/model"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {
	log.Print("[OrderRepository]...Migrate")
	db.AutoMigrate(&model.TCustomer{})
	db.AutoMigrate(&model.TProduct{})

	products := []model.TProduct{
		{
			Id:       1,
			Name:     "Pecel Lele",
			Quantity: 100,
			Price:    1000,
		},
		{
			Id:       2,
			Name:     "Baso Sapi",
			Quantity: 100,
			Price:    1000,
		},
		{
			Id:       3,
			Name:     "Batagor",
			Quantity: 100,
			Price:    1000,
		},
	}

	db.Create(&products)

	db.AutoMigrate(&model.TOrder{})
	return db.AutoMigrate(&model.TOrderItem{})
}
