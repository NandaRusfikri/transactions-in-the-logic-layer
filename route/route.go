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
	orderService := service.NewOrderService(orderRepository, db, uw)

	orderController := controller.NewOrderController(orderService)

	orders := httpRouter.Group("order")

	orders.POST("/", orderController.AddOrder)

	//httpRouter.POST("/money-transfer", middleware.DBTransactionMiddleware(db), orderController.TransferMoney)
	httpRouter.Run()

}
