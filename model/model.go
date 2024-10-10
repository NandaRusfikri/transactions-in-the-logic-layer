package model

type OrderRequest struct {
	CustomerId int                `json:"customer_id" binding:"required"`
	Items      []OrderItemRequest `json:"items"`
}
type OrderItemRequest struct {
	ProductId uint   `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Note      string `json:"note"`
}
