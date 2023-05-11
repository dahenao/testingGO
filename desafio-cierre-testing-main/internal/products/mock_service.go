package products

import "github.com/stretchr/testify/mock"

type serviceMock struct {
	mock.Mock
}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) GetAllBySeller(sellerID string) ([]Product, error) {
	// input
	// -> for expectations (to check if the input is the same as the expected)

	args := s.Called(sellerID)

	return args.Get(0).([]Product), args.Error(1)
}
