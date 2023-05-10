package prey

import (
	"integrationtests/pkg/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeed(t *testing.T) {
	//arrange
	Data := map[string]interface{}{
		"tuna_speed": 116.9,
	}

	storageMock := storage.NewMockStorage()
	storageMock.Data = Data
	Prey := CreateTuna(storageMock)
	expectedSpeed := 116.9
	//act
	Speed := Prey.GetSpeed()

	//assert
	assert.Equal(t, expectedSpeed, Speed)

}
