package transaction

import (
	"context"
	"github.com/stretchr/testify/mock"
)

// MockUoW is a mock of the UoW struct
type MockUoW struct {
	mock.Mock
}

// Mock the WithTx method
func (m *MockUoW) WithTx(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	args := m.Called(ctx, fn)
	return args.Get(0), args.Error(1)
}
