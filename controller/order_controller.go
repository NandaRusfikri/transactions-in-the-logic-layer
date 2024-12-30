package controller

import (
	"go-transaction/model"
	"go-transaction/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderController : represent the order's controller contract
type OrderController interface {
	Order(*gin.Context)
	Orders(c *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

// NewOrderController -> returns new order controller
func NewOrderController(s service.OrderService) OrderController {
	return orderController{
		orderService: s,
	}
}

func (u orderController) Order(c *gin.Context) {
	var order model.OrderRequest
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.orderService.Order(c, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
func (u orderController) Orders(c *gin.Context) {

	res, err := u.orderService.Orders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error get orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
