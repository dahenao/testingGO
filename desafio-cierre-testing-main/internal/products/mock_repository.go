package products

import "github.com/stretchr/testify/mock"

type repositoryMock struct {
	mock.Mock
}

func NewRepositoryMock() *repositoryMock { //retorna puntero para poder tener control de la struct
	return &repositoryMock{}
}

func (r *repositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	// input
	// -> for expectations (to check if the input is the same as the expected)

	args := r.Called(sellerID)

	return args.Get(0).([]Product), args.Error(1)
}
