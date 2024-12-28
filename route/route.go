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

	if err := repository.Migrate(db); err != nil {
		log.Fatal("Order migrate err", err)
	}

	orderRepository := repository.NewOrderRepository(db)
	productRepository := repository.NewProductRepository(db)
	uw := transaction.NewUW(db)

	orderService := service.NewOrderService(orderRepository, productRepository, uw)
	orderController := controller.NewOrderController(orderService)

	v1 := httpRouter.Group("v1")
	v1.POST("order", orderController.Order)
	v1.GET("orders", orderController.Orders)

	httpRouter.Run(":9999")

}
