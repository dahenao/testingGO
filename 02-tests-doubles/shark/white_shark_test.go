package shark

import (
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunt(t *testing.T) {

	t.Run("successffully hunt", func(t *testing.T) {
		//assert
		simulatorMock := simulator.CreateMockSimulator()
		simulatorMock.CanCatchMock = true
		simulatorMock.SpyGetLinearDistance = false
		simulatorMock.SpycanCath = false
		whiteShark := CreateWhiteShark(simulatorMock)
		prey := prey.NewPreyStub()
		prey.MaxSpeed = 252

		//act
		result := whiteShark.Hunt(prey)

		//assert
		assert.Nil(t, result, "result of error not is nil")
		assert.True(t, simulatorMock.SpyGetLinearDistance, "GetLinearDistance must be true")

	})

	t.Run("shark not hunt slow shark", func(t *testing.T) {
		//assert
		simulatorMock := simulator.CreateMockSimulator()
		simulatorMock.CanCatchMock = false
		simulatorMock.SpyGetLinearDistance = false
		simulatorMock.SpycanCath = false
		whiteShark := CreateWhiteShark(simulatorMock)
		prey := prey.NewPreyStub()
		prey.MaxSpeed = 252

		//act
		result := whiteShark.Hunt(prey)

		//assert
		//assert.ErrorIs(t, result, fmt.Errorf("could not hunt the prey"), "incorrect error")
		assert.Equal(t, result.Error(), "could not hunt the prey", "incorrect error")
		assert.True(t, simulatorMock.SpyGetLinearDistance, "GetLinearDistance must be true")

	})

	t.Run("shark not hunt by large distance", func(t *testing.T) {
		//assert
		simulatorMock := simulator.CreateMockSimulator()
		simulatorMock.CanCatchMock = false
		simulatorMock.SpyGetLinearDistance = false
		simulatorMock.SpycanCath = false
		whiteShark := CreateWhiteShark(simulatorMock)
		prey := prey.NewPreyStub()
		prey.MaxSpeed = 252

		//act
		result := whiteShark.Hunt(prey)

		//assert
		//assert.ErrorIs(t, result, fmt.Errorf("could not hunt the prey"), "incorrect error")
		assert.Equal(t, result.Error(), "could not hunt the prey", "incorrect error")
		assert.True(t, simulatorMock.SpyGetLinearDistance, "GetLinearDistance must be true")

	})
}
