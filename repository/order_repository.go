package repository

import (
	"context"
	"go-transaction/model"
	"go-transaction/transaction"
	"log"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

type OrderRepository interface {
	AddOrder(ctx context.Context, data model.TableOrder) (model.TableOrder, error)
	AddOrderItem(ctx context.Context, item model.TableOrderItem) (model.TableOrderItem, error)
	UpdateStock(ctx context.Context, productId, quantity uint) (model.TableProduct, error)
	GetAll() ([]model.TableOrder, error)
	WithTrx(*gorm.DB) orderRepository
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
	u.DB.AutoMigrate(&model.TableProduct{})

	products := []model.TableProduct{
		{
			Id:       1,
			Name:     "Sepatu",
			Quantity: 100,
		},
		{
			Id:       2,
			Name:     "Baju",
			Quantity: 100,
		},
		{
			Id:       3,
			Name:     "Celana",
			Quantity: 100,
		},
	}

	u.DB.Create(&products)

	u.DB.AutoMigrate(&model.TableOrder{})
	return u.DB.AutoMigrate(&model.TableOrderItem{})
}

func (u orderRepository) AddOrder(ctx context.Context, data model.TableOrder) (model.TableOrder, error) {
	log.Print("[OrderRepository]...AddOrder")

	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}
	err := tx.Create(&data).Error
	return data, err
}
func (u orderRepository) UpdateStock(ctx context.Context, productId, quantity uint) (model.TableProduct, error) {
	log.Print("[OrderRepository]...UpdateStock")

	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}

	var data model.TableProduct

	err := u.DB.Where("id = ? ", productId).First(&data)
	if err.Error != nil {
		return data, err.Error
	}

	update := model.TableProduct{
		Id:       productId,
		Quantity: data.Quantity - int(quantity),
	}
	errs := tx.Updates(&update).Error
	return data, errs
}

func (u orderRepository) AddOrderItem(ctx context.Context, data model.TableOrderItem) (model.TableOrderItem, error) {
	log.Print("[OrderRepository]...Activity")
	tx, ok := transaction.GetTx(ctx)
	if !ok {
		tx = u.DB
	}
	err := tx.Create(&data).Error
	return data, err

}

func (u orderRepository) GetAll() (users []model.TableOrder, err error) {
	log.Print("[OrderRepository]...Get All")
	err = u.DB.Find(&users).Error
	return users, err

}

func (u orderRepository) WithTrx(trxHandle *gorm.DB) orderRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}
