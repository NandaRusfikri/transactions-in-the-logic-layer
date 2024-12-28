package service

import (
	"context"
	"go-transaction/model"
	"go-transaction/repository"
	"go-transaction/transaction"
	"log"
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

func (u orderService) Order(ctx2 context.Context, param model.OrderRequest) (model.TOrder, error) {
	log.Print("[orderService]...add Order")

	v, err := u.UW.WithTx(ctx2, func(ctx context.Context) (interface{}, error) {
		dataOrder, err := u.orderRepository.AddOrder(ctx, model.TOrder{CustomerId: param.CustomerId})
		if err != nil {
			return model.TOrder{}, err
		}

		for _, item := range param.Items {
			product, err := u.productRepository.GetProduct(ctx, model.TProduct{Id: item.ProductId})
			if err != nil {
				return nil, err
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

			_, err = u.orderRepository.UpdateStock(ctx, item.ProductId, item.Quantity)
			if err != nil {
				return model.TOrder{}, err
			}

		}

		return dataOrder, nil

	})
	if err != nil {
		return model.TOrder{}, err
	}

	return v.(model.TOrder), err

}

func (u orderService) Orders() ([]model.TOrder, error) {
	log.Print("[orderService]...Orders")

	return u.orderRepository.GetAll()

}
