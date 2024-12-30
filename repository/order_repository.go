package repository

import (
	"context"
	"go-transaction/model"
	"go-transaction/transaction"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepository struct {
	DB *gorm.DB
}

type OrderRepository interface {
	AddOrder(ctx context.Context, data model.TOrder) (model.TOrder, error)
	AddOrderItem(ctx context.Context, item model.TOrderItem) (model.TOrderItem, error)
	GetAll() ([]model.TOrder, error)
}

// NewOrderRepository -> returns new user repository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return orderRepository{
		DB: db,
	}
}

func (u orderRepository) AddOrder(ctx context.Context, data model.TOrder) (model.TOrder, error) {

	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}
	err := tx.Create(&data).Error
	return data, err
}

func (u orderRepository) AddOrderItem(ctx context.Context, data model.TOrderItem) (model.TOrderItem, error) {
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}
	err := tx.Create(&data).Error
	return data, err

}

func (u orderRepository) GetAll() (data []model.TOrder, err error) {
	err = u.DB.Preload(clause.Associations).Find(&data).Error
	return data, err

}
