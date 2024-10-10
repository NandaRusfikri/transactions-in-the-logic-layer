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
	orderRepository repository.OrderRepository
	UW              transaction.UoW
}

func NewOrderService(r repository.OrderRepository, uw transaction.UoW) OrderService {
	return orderService{
		orderRepository: r,
		UW:              uw,
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
			_, err = u.orderRepository.AddOrderItem(ctx, model.TOrderItem{
				OrderId:   dataOrder.Id,
				Quantity:  item.Quantity,
				ProductId: item.ProductId,
				Note:      item.Note,
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
