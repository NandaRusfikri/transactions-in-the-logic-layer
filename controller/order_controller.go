package controller

import (
	"go-transaction/model"
	"go-transaction/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderController : represent the order's controller contract
type OrderController interface {
	AddOrder(*gin.Context)
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

func (u orderController) AddOrder(c *gin.Context) {
	log.Print("[OrderController]...add TableOrder")
	var order model.CreateOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.orderService.CreateOrder(c, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
