package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCatch(t *testing.T) {

	t.Run("distance is big", func(t *testing.T) {
		//arrange
		simu := NewCatchSimulator(5)
		expectResult := false

		//act

		result := simu.CanCatch(500.0, 122.0, 100.0)

		//assert
		assert.False(t, result, expectResult)
	})

	t.Run("shark is slow", func(t *testing.T) {
		//arrange
		simu := NewCatchSimulator(5)
		expectResult := false

		//act

		result := simu.CanCatch(10.0, 99.0, 100.0)

		//assert
		assert.False(t, result, expectResult)
	})

	t.Run("shark fast and hunt", func(t *testing.T) {
		//arrange
		simu := NewCatchSimulator(5)
		expectResult := false

		//act

		result := simu.CanCatch(10.0, 299.0, 100.0)

		//assert
		assert.True(t, result, expectResult)
	})

}

func TestGetLinearDistance(t *testing.T) {

	t.Run("distance between 50, 100", func(t *testing.T) {
		//arrange
		simu := NewCatchSimulator(5)
		expectResult := 111.80339887498948

		//act

		result := simu.GetLinearDistance([2]float64{50.0, 100.0})

		//assert

		assert.Equal(t, expectResult, result)
	})
}
