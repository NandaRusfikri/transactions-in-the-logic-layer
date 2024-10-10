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

func mockUnitOfWork(mockUW *transaction.MockUoW, order model.TableOrder, err error) {
	mockUW.On("WithTx", mock.Anything, mock.AnythingOfType("func(context.Context) (interface{}, error)")).
		Return(order, err).Once()
}
func createInput() model.CreateOrder {
	return model.CreateOrder{
		CustomerId: 11,
		Items: []model.CreateOrderItem{
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
	service := NewOrderService(mockRepo, gormDB, uW)

	var errorDB = errors.New("error db")

	inputUsecase := createInput()
	order := model.TableOrder{
		CustomerId: 11,
	}

	tests := []struct {
		name           string
		input          model.CreateOrder
		mock           func()
		expectedErr    error
		expectedResult model.TableOrder
	}{
		{
			name:  "1.success new order",
			input: inputUsecase,
			mock: func() {
				orderRes := model.TableOrder{ID: 5}
				mockRepo.On("AddOrder", mock.Anything, order).Return(orderRes, nil).Once()
				for _, item := range inputUsecase.Items {
					mockRepo.On("AddOrderItem", mock.Anything, model.TableOrderItem{
						OrderId:   orderRes.ID,
						ProductId: item.ProductId,
						Quantity:  item.Quantity,
						Note:      item.Note,
					}).Return(model.TableOrderItem{}, nil).Once()
					mockRepo.On("UpdateStock", mock.Anything, item.ProductId, item.Quantity).Return(model.TableProduct{}, nil).Once()
				}
				mockUnitOfWork(mockUW, model.TableOrder{ID: 5}, nil)
			},
			expectedErr:    nil,
			expectedResult: model.TableOrder{ID: 5},
		},
		{
			name:  "2.failed call method AddOrder",
			input: inputUsecase,
			mock: func() {
				mockRepo.On("AddOrder", mock.Anything, order).Return(model.TableOrder{}, errorDB).Once()
				mockUnitOfWork(mockUW, model.TableOrder{}, errorDB)
			},
			expectedErr: errorDB,
		},
		{
			name:  "3.failed call method AddOrderItem",
			input: inputUsecase,
			mock: func() {
				mockRepo.On("AddOrder", mock.Anything, order).Return(model.TableOrder{}, nil).Once()
				for _, item := range inputUsecase.Items {
					mockRepo.On("AddOrderItem", mock.Anything, model.TableOrderItem{
						OrderId:   order.ID,
						ProductId: item.ProductId,
						Quantity:  item.Quantity,
						Note:      item.Note,
					}).Return(model.TableOrderItem{}, errorDB).Once()
				}
				mockUnitOfWork(mockUW, model.TableOrder{}, errorDB)
			},
			expectedErr: errorDB,
		},
		{
			name:  "4.failed call method UpdateStock",
			input: inputUsecase,
			mock: func() {
				mockRepo.On("AddOrder", mock.Anything, order).Return(model.TableOrder{}, nil).Once()
				for _, item := range inputUsecase.Items {
					mockRepo.On("AddOrderItem", mock.Anything, model.TableOrderItem{
						OrderId:   order.ID,
						ProductId: item.ProductId,
						Quantity:  item.Quantity,
						Note:      item.Note,
					}).Return(model.TableOrderItem{}, nil).Once()
					mockRepo.On("UpdateStock", mock.Anything, item.ProductId, item.Quantity).Return(model.TableProduct{}, errorDB).Once()
				}
				mockUnitOfWork(mockUW, model.TableOrder{}, errorDB)
			},
			expectedErr: errorDB,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			result, err := service.CreateOrder(context.TODO(), inputUsecase)
			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}

}
