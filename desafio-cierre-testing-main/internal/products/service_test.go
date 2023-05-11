package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller(t *testing.T) {
	//arrange
	rp := NewRepositoryMock()
	sv := NewService(rp)
	rp.On("GetAllBySeller", "FEX112AC").
		Return([]Product{{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		}}, nil)
	expectedPrds := []Product{{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	}}
	//act
	prds, err := sv.GetAllBySeller("FEX112AC")

	//assert
	assert.NoError(t, err)
	assert.Equal(t, expectedPrds, prds)

}
