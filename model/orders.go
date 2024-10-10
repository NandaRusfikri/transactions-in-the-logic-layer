package model

import (
	"time"
)

type (
	TableOrder struct {
		ID         uint `gorm:"primarykey"`
		CreatedAt  time.Time
		CustomerId int `gorm:"column:customer_id" json:"customer_id"`
	}
	TableOrderItem struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		OrderId   uint   `gorm:"column:order_id"`
		ProductId uint   `gorm:"column:product_id"`
		Quantity  uint   `gorm:"column:quantity"`
		Note      string `gorm:"type:varchar(20)"`
	}
	TableProduct struct {
		Id       uint   `gorm:"column:id;primary_key"`
		Name     string `gorm:"column:name"`
		Quantity int    `gorm:"column:quantity"`
	}
)

func (TableOrder) TableName() string {
	return "orders"
}
func (TableProduct) TableName() string {
	return "products"
}
func (TableOrderItem) TableName() string {
	return "order_items"
}

type CreateOrder struct {
	CustomerId int               `json:"customer_id"`
	Items      []CreateOrderItem `json:"items"`
}
type CreateOrderItem struct {
	ProductId uint   `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Note      string `json:"note"`
}
