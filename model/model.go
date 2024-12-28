package model

type OrderRequest struct {
	CustomerId uint64             `json:"customer_id" binding:"required"`
	Items      []OrderItemRequest `json:"items"`
}
type OrderItemRequest struct {
	ProductId uint64 `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Note      string `json:"note"`
}
