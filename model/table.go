package model

import (
	"time"
)

type (
	TOrder struct {
		Id         uint         `gorm:"primarykey" json:"id"`
		CreatedAt  time.Time    `json:"created_at"`
		CustomerId int          `gorm:"column:customer_id" json:"customer_id"`
		Items      []TOrderItem `gorm:"foreignKey:OrderId;references:Id"`
	}
	TOrderItem struct {
		Id        uint      `gorm:"primarykey" json:"id"`
		CreatedAt time.Time `json:"-"`
		OrderId   uint      `gorm:"column:order_id" json:"order_id"`
		ProductId uint      `gorm:"column:product_id" json:"product_id"`
		Quantity  uint      `gorm:"column:quantity" json:"quantity"`
		Note      string    `gorm:"type:varchar(20)" json:"note"`
	}
	TProduct struct {
		Id       uint   `gorm:"column:id;primary_key"`
		Name     string `gorm:"column:name"`
		Quantity int    `gorm:"column:quantity"`
	}
)

func (TOrder) TableName() string {
	return "orders"
}
func (TProduct) TableName() string {
	return "products"
}
func (TOrderItem) TableName() string {
	return "order_items"
}
