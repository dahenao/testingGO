package shark

import (
	"integrationtests/pkg/storage"
	"integrationtests/prey"
	"integrationtests/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunt(t *testing.T) {

	t.Run("shark hunt succesfully", func(t *testing.T) {
		//arrange

		Data := map[string]interface{}{
			"white_shark_speed": 244.2,
			"white_shark_x":     1.0,
			"white_shark_y":     2.0,
			"tuna_speed":        116.9,
		}

		storageMock := storage.NewMockStorage()
		storageMock.Data = Data
		simu := simulator.NewCatchSimulator(3.0)
		whiteShark := CreateWhiteShark(simu, storageMock)
		prey := prey.CreateTuna(storageMock)

		//act
		result := whiteShark.Hunt(prey)

		//assert
		assert.Nil(t, result)

	})

	t.Run("shark not hunt by be slow", func(t *testing.T) {
		//arrange

		Data := map[string]interface{}{
			"white_shark_speed": 44.2,
			"white_shark_x":     1.0,
			"white_shark_y":     2.0,
			"tuna_speed":        116.9,
		}

		storageMock := storage.NewMockStorage()
		storageMock.Data = Data
		simu := simulator.NewCatchSimulator(3.0)
		whiteShark := CreateWhiteShark(simu, storageMock)
		prey := prey.CreateTuna(storageMock)

		//act
		result := whiteShark.Hunt(prey)

		//assert
		assert.NotNil(t, result)
		assert.Error(t, result, "could not hunt the prey")

	})

	t.Run("shark not hunt by be slow", func(t *testing.T) {
		//arrange

		Data := map[string]interface{}{
			"white_shark_speed": 244.2,
			"white_shark_x":     488.0,
			"white_shark_y":     499.0,
			"tuna_speed":        116.9,
		}

		storageMock := storage.NewMockStorage()
		storageMock.Data = Data
		simu := simulator.NewCatchSimulator(3.0)
		whiteShark := CreateWhiteShark(simu, storageMock)
		prey := prey.CreateTuna(storageMock)

		//act
		result := whiteShark.Hunt(prey)

		//assert
		assert.NotNil(t, result)
		assert.Error(t, result, "could not hunt the prey")

	})

}
