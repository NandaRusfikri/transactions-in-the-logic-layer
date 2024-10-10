package route

import (
	"go-transaction/controller"
	"go-transaction/repository"
	"go-transaction/service"
	"go-transaction/transaction"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	orderRepository := repository.NewOrderRepository(db)
	uw := transaction.NewUW(db)

	if err := orderRepository.Migrate(); err != nil {
		log.Fatal("Order migrate err", err)
	}
	orderService := service.NewOrderService(orderRepository, uw)
	orderController := controller.NewOrderController(orderService)

	v1 := httpRouter.Group("v1")
	v1.POST("order", orderController.Order)
	v1.GET("orders", orderController.Orders)

	httpRouter.Run(":9999")

}
