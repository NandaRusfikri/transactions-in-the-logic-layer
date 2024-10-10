
# Implement Database Transactions in Logic Layer (GORM + GIN)

## API Documentation

### POST v1/order

This endpoint is used to create a new order for a customer. The request body should include customer details and the list of items in the order.

### Request Body

The request body should be in JSON format and should contain the following structure:

```json
{
  "customer_id": 1,
  "items": [
    {
      "product_id": 1,
      "quantity": 2,
      "note": "No onions, please"
    },
    {
      "product_id": 2,
      "quantity": 1,
      "note": ""
    }
  ]
}
``` 

```shell
curl -X POST http://localhost:9999/v1/order \
-H "Content-Type: application/json" \
-d '{
  "customer_id": 1,
  "items": [
    {
      "product_id": 1,
      "quantity": 2,
      "note": "No onions, please"
    },
    {
      "product_id": 2,
      "quantity": 1,
      "note": ""
    }
  ]
}'

```