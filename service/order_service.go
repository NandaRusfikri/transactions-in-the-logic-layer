package service

import (
	"context"
	"go-transaction/model"
	"go-transaction/repository"
	"go-transaction/transaction"
	"log"

	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order model.CreateOrder) (model.TableOrder, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
	db              *gorm.DB
	UW              transaction.UoW
}

func NewOrderService(r repository.OrderRepository, db *gorm.DB, uw transaction.UoW) OrderService {
	return orderService{
		orderRepository: r,
		db:              db,
		UW:              uw,
	}
}

func (u orderService) CreateOrder(ctx2 context.Context, param model.CreateOrder) (model.TableOrder, error) {
	log.Print("[orderService]...add TableOrder")

	v, err := u.UW.WithTx(ctx2, func(ctx context.Context) (interface{}, error) {
		dataOrder, err := u.orderRepository.AddOrder(ctx, model.TableOrder{CustomerId: param.CustomerId})
		if err != nil {
			return model.TableOrder{}, err
		}

		for _, item := range param.Items {
			_, err = u.orderRepository.AddOrderItem(ctx, model.TableOrderItem{
				OrderId:   dataOrder.ID,
				Quantity:  item.Quantity,
				ProductId: item.ProductId,
				Note:      item.Note,
			})
			if err != nil {
				return model.TableOrder{}, err
			}

			_, err = u.orderRepository.UpdateStock(ctx, item.ProductId, item.Quantity)
			if err != nil {
				return model.TableOrder{}, err
			}

		}

		return dataOrder, nil

	})
	if err != nil {
		return model.TableOrder{}, err
	}

	return v.(model.TableOrder), err

}
