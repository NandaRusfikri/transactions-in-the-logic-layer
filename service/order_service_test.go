package service

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-transaction/model"
	"go-transaction/repository"
	"go-transaction/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func mockUnitOfWork(mockUW *transaction.MockUoW, order model.TOrder, err error) {
	mockUW.On("WithTx", mock.Anything, mock.AnythingOfType("func(context.Context) (interface{}, error)")).
		Return(order, err).Once()
}
func createInput() model.OrderRequest {
	return model.OrderRequest{
		CustomerId: 11,
		Items: []model.OrderItemRequest{
			{
				Quantity:  1,
				ProductId: 1,
			},
			{
				Quantity:  2,
				ProductId: 2,
			},
		},
	}
}

func TestAddOrder(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	mockRepo := new(repository.MockOrderRepository)
	mockUW := new(transaction.MockUoW)
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	uW := transaction.NewUW(gormDB)
	service := NewOrderService(mockRepo, uW)

	var errorDB = errors.New("error db")

	inputUsecase := createInput()
	order := model.TOrder{
		CustomerId: 11,
	}

	tests := []struct {
		name           string
		input          model.OrderRequest
		mock           func()
		expectedErr    error
		expectedResult model.TOrder
	}{
		{
			name:  "1.success new order",
			input: inputUsecase,
			mock: func() {
				orderRes := model.TOrder{Id: 5}
				mockRepo.On("AddOrder", mock.Anything, order).Return(orderRes, nil).Once()
				for _, item := range inputUsecase.Items {
					mockRepo.On("AddOrderItem", mock.Anything, model.TOrderItem{
						OrderId:   orderRes.Id,
						ProductId: item.ProductId,
						Quantity:  item.Quantity,
						Note:      item.Note,
					}).Return(model.TOrderItem{}, nil).Once()
					mockRepo.On("UpdateStock", mock.Anything, item.ProductId, item.Quantity).Return(model.TProduct{}, nil).Once()
				}
				mockUnitOfWork(mockUW, model.TOrder{Id: 5}, nil)
			},
			expectedErr:    nil,
			expectedResult: model.TOrder{Id: 5},
		},
		{
			name:  "2.failed call method Order",
			input: inputUsecase,
			mock: func() {
				mockRepo.On("AddOrder", mock.Anything, order).Return(model.TOrder{}, errorDB).Once()
				mockUnitOfWork(mockUW, model.TOrder{}, errorDB)
			},
			expectedErr: errorDB,
		},
		{
			name:  "3.failed call method AddOrderItem",
			input: inputUsecase,
			mock: func() {
				mockRepo.On("AddOrder", mock.Anything, order).Return(model.TOrder{}, nil).Once()
				for _, item := range inputUsecase.Items {
					mockRepo.On("AddOrderItem", mock.Anything, model.TOrderItem{
						OrderId:   order.Id,
						ProductId: item.ProductId,
						Quantity:  item.Quantity,
						Note:      item.Note,
					}).Return(model.TOrderItem{}, errorDB).Once()
				}
				mockUnitOfWork(mockUW, model.TOrder{}, errorDB)
			},
			expectedErr: errorDB,
		},
		{
			name:  "4.failed call method UpdateStock",
			input: inputUsecase,
			mock: func() {
				mockRepo.On("AddOrder", mock.Anything, order).Return(model.TOrder{}, nil).Once()
				for _, item := range inputUsecase.Items {
					mockRepo.On("AddOrderItem", mock.Anything, model.TOrderItem{
						OrderId:   order.Id,
						ProductId: item.ProductId,
						Quantity:  item.Quantity,
						Note:      item.Note,
					}).Return(model.TOrderItem{}, nil).Once()
					mockRepo.On("UpdateStock", mock.Anything, item.ProductId, item.Quantity).Return(model.TProduct{}, errorDB).Once()
				}
				mockUnitOfWork(mockUW, model.TOrder{}, errorDB)
			},
			expectedErr: errorDB,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			result, err := service.Order(context.TODO(), inputUsecase)
			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}

}
