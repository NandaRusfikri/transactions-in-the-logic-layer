package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-transaction/model"
	"go-transaction/repository"
	"go-transaction/transaction"
	"log"
	"time"
)

type OrderService interface {
	Order(ctx context.Context, order model.OrderRequest) (model.TOrder, error)
	Orders() ([]model.TOrder, error)
}

type orderService struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
	UW                transaction.UoW
}

func NewOrderService(order repository.OrderRepository, product repository.ProductRepository, uw transaction.UoW) OrderService {
	return orderService{
		orderRepository:   order,
		productRepository: product,
		UW:                uw,
	}
}

func (u orderService) Order(ctx context.Context, param model.OrderRequest) (model.TOrder, error) {

	key := "order_counter:" + time.Now().Format("2006-01-02")
	counter, err := model.RedisClient.Incr(ctx, key).Result()
	if err != nil {
		log.Fatalf("Failed to increment order number: %v", err)
	}

	day := time.Now().Format("2006-01-02")
	orderNumber := fmt.Sprintf("INV/%v/%v", day, counter)
	v, err := u.UW.WithTx(ctx, func(ctx context.Context) (interface{}, error) {
		dataOrder, err := u.orderRepository.AddOrder(ctx, model.TOrder{CustomerId: param.CustomerId, OrderNumber: orderNumber})
		if err != nil {
			return model.TOrder{}, err
		}

		for _, item := range param.Items {
			ctxLocking := context.WithValue(ctx, transaction.CtxLocking, "UPDATE")
			product, err := u.productRepository.GetProduct(ctxLocking, model.TProduct{Id: item.ProductId})
			if err != nil {
				return model.TOrder{}, err
			}
			if (product.Quantity - item.Quantity) <= 1 {
				return model.TOrder{}, fmt.Errorf("quantity kureng")
			}

			_, err = u.orderRepository.AddOrderItem(ctx, model.TOrderItem{
				OrderId:      dataOrder.Id,
				ProductId:    product.Id,
				ProductName:  product.Name,
				ProductPrice: product.Price,
				Quantity:     item.Quantity,
				Note:         item.Note,
			})
			if err != nil {
				return model.TOrder{}, err
			}

			product.Quantity = product.Quantity - item.Quantity
			_, err = u.productRepository.Updates(ctx, product)
			if err != nil {
				return model.TOrder{}, err
			}

		}

		return dataOrder, nil

	})
	if err != nil {
		logrus.Errorf("[orderService] err:%v req:%v \n", err, param.Items)
		return model.TOrder{}, err
	}

	return v.(model.TOrder), err

}

func (u orderService) Orders() ([]model.TOrder, error) {
	log.Print("[orderService]...Orders")

	return u.orderRepository.GetAll()

}
