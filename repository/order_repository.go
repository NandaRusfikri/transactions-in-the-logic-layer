package repository

import (
	"context"
	"go-transaction/model"
	"go-transaction/transaction"
	"gorm.io/gorm/clause"
	"log"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

type OrderRepository interface {
	AddOrder(ctx context.Context, data model.TOrder) (model.TOrder, error)
	AddOrderItem(ctx context.Context, item model.TOrderItem) (model.TOrderItem, error)
	UpdateStock(ctx context.Context, productId, quantity uint) (model.TProduct, error)
	GetAll() ([]model.TOrder, error)
	Migrate() error
}

// NewOrderRepository -> returns new user repository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return orderRepository{
		DB: db,
	}
}

func (u orderRepository) Migrate() error {
	log.Print("[OrderRepository]...Migrate")
	u.DB.AutoMigrate(&model.TProduct{})

	products := []model.TProduct{
		{
			Id:       1,
			Name:     "Pecel Lele",
			Quantity: 100,
		},
		{
			Id:       2,
			Name:     "Baso Sapi",
			Quantity: 100,
		},
		{
			Id:       3,
			Name:     "Batagor",
			Quantity: 100,
		},
	}

	u.DB.Create(&products)

	u.DB.AutoMigrate(&model.TOrder{})
	return u.DB.AutoMigrate(&model.TOrderItem{})
}

func (u orderRepository) AddOrder(ctx context.Context, data model.TOrder) (model.TOrder, error) {
	log.Print("[OrderRepository]...Order")

	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}
	err := tx.Create(&data).Error
	return data, err
}
func (u orderRepository) UpdateStock(ctx context.Context, productId, quantity uint) (model.TProduct, error) {
	log.Print("[OrderRepository]...UpdateStock")

	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}

	var data model.TProduct

	err := u.DB.Where("id = ? ", productId).First(&data)
	if err.Error != nil {
		return data, err.Error
	}

	update := model.TProduct{
		Id:       productId,
		Quantity: data.Quantity - int(quantity),
	}
	errs := tx.Updates(&update).Error
	return data, errs
}

func (u orderRepository) AddOrderItem(ctx context.Context, data model.TOrderItem) (model.TOrderItem, error) {
	log.Print("[OrderRepository]...Activity")
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}
	err := tx.Create(&data).Error
	return data, err

}

func (u orderRepository) GetAll() (data []model.TOrder, err error) {
	log.Print("[OrderRepository]...Get All")
	err = u.DB.Preload(clause.Associations).Find(&data).Error
	return data, err

}
