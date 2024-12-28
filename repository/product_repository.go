package repository

import (
	"context"
	"go-transaction/model"
	"log"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	GetProduct(ctx context.Context, data model.TProduct) (model.TProduct, error)
}

// NewproductRepository -> returns new user repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{
		DB: db,
	}
}

func (u productRepository) GetProduct(ctx context.Context, data model.TProduct) (model.TProduct, error) {
	log.Print("[productRepository]...product")

	db := u.DB.Model(model.TProduct{})

	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	err := db.First(&data).Error
	return data, err
}
