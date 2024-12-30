package repository

import (
	"context"
	"go-transaction/model"
	"go-transaction/transaction"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	GetProduct(ctx context.Context, data model.TProduct) (model.TProduct, error)
	Updates(ctx context.Context, data model.TProduct) (model.TProduct, error)
}

// NewproductRepository -> returns new user repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{
		DB: db,
	}
}

func (u productRepository) GetProduct(ctx context.Context, data model.TProduct) (model.TProduct, error) {

	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}

	strength, ok := transaction.GetLocking(ctx)
	if ok {
		tx = tx.Clauses(clause.Locking{Strength: strength})
	}

	if data.Id != 0 {
		tx = tx.Where("id = ?", data.Id)
	}
	err := tx.First(&data).Error
	return data, err
}

func (u productRepository) Updates(ctx context.Context, data model.TProduct) (model.TProduct, error) {

	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}
	err := tx.Updates(&data).Error
	return data, err
}
