package model

type OrderRequest struct {
	IdempotencyKey string             `json:"idempotency_key"`
	CustomerId     uint64             `json:"customer_id" binding:"required"`
	Items          []OrderItemRequest `json:"items"`
}
type OrderItemRequest struct {
	ProductId uint64 `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Note      string `json:"note"`
}
