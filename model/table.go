package model

import (
	"time"
)

type (
	TCustomer struct {
		Id        uint64    `gorm:"primarykey" json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Name      string    `gorm:"column:name" json:"name"`
		Email     string    `gorm:"column:email" json:"email"`
		Phone     string    `gorm:"column:phone" json:"phone"`
		Address   string    `gorm:"column:address" json:"address"`
		City      string    `gorm:"column:city" json:"city"`
	}
	TOrder struct {
		Id         uint64       `gorm:"primarykey" json:"id"`
		CreatedAt  time.Time    `json:"created_at"`
		CustomerId uint64       `gorm:"column:customer_id" json:"customer_id"`
		Items      []TOrderItem `gorm:"foreignKey:OrderId;references:Id"`
	}
	TOrderItem struct {
		Id           uint64    `gorm:"primarykey" json:"id"`
		CreatedAt    time.Time `json:"-"`
		OrderId      uint64    `gorm:"column:order_id" json:"order_id"`
		ProductId    uint64    `gorm:"column:product_id" json:"product_id"`
		ProductName  string    `gorm:"column:product_name;type:varchar(100)" json:"product_name"`
		ProductPrice uint      `gorm:"column:product_price" json:"product_price"` // Harga produk saat order
		Quantity     uint      `gorm:"column:quantity" json:"quantity"`
		Note         string    `gorm:"type:varchar(20)" json:"note"`
	}
	TProduct struct {
		Id       uint64 `gorm:"column:id;primary_key"`
		Name     string `gorm:"column:name"`
		Quantity int    `gorm:"column:quantity"`
		Price    uint   `gorm:"column:price"`
	}
)

func (TCustomer) TableName() string {
	return "customers"
}
func (TOrder) TableName() string {
	return "orders"
}
func (TProduct) TableName() string {
	return "products"
}
func (TOrderItem) TableName() string {
	return "order_items"
}
