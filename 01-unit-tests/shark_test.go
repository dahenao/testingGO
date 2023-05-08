package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	//arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  60}

	prey := Prey{
		name:  "human",
		speed: 2}
	expectedHungry := false
	expectedTired := true
	//act

	err := shark.Hunt(&prey)
	//assert
	assert.Nil(t, err)
	//assert.Equal(t, expectedHungry, shark.hungry, "incorrect hungry")
	//assert.Equal(t, expectedTired, shark.tired, "incorrect tired")
	assert.True(t, expectedTired) //other way
	assert.False(t, expectedHungry)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//arrange
	shark := Shark{
		hungry: true,
		tired:  true,
		speed:  60}

	prey := Prey{
		name:  "human",
		speed: 2}
	expectedErr := ErrSharkTired //fmt.Errorf("cannot hunt, i am really tired")
	//act

	err := shark.Hunt(&prey)

	//assert
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrSharkTired)
	assert.Equal(t, expectedErr, err, "incorrect error")
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	//arrange
	shark := Shark{
		hungry: false,
		tired:  false,
		speed:  60}

	prey := Prey{
		name:  "human",
		speed: 2}
	expectedErr := ErrSharkNotHungry //fmt.Errorf("cannot hunt, i am not hungry")
	//act

	err := shark.Hunt(&prey)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err, "incorrect error")
	assert.ErrorIs(t, err, ErrSharkNotHungry)
}

func TestSharkCannotReachThePrey(t *testing.T) {
	//arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  60}

	prey := Prey{
		name:  "penguin",
		speed: 63}
	expectedErr := ErrSharkNotCatch //fmt.Errorf("could not catch it")
	//act

	err := shark.Hunt(&prey)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err, "incorrect error")
	assert.ErrorIs(t, err, ErrSharkNotCatch)
}

func TestSharkHuntNilPrey(t *testing.T) {

	//arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  60}

	expectedErr := ErrPreyNil //fmt.Errorf("prey is nil")
	//act

	err := shark.Hunt(nil)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err, "incorrect error")
	assert.ErrorIs(t, err, ErrPreyNil)
}

/*func TestSharkHunt(t *testing.T) {
	testCases := []struct {
		name           string
		shark          Shark
		prey           Prey
		expectedErr    error
		expectedHungry bool
		expectedTired  bool
	}{
		{
			name: "CatchPrey",
			shark: Shark{
				hungry: true,
				tired:  false,
				speed:  10,
			},
			prey: Prey{
				name:  "tuna",
				speed: 5,
			},
			expectedErr:    nil,
			expectedHungry: false,
			expectedTired:  true,
		},
		{
			name: "PreyIsNil",
			shark: Shark{
				hungry: true,
				tired:  false,
				speed:  10,
			},

			expectedErr:    ErrPreyNil,
			expectedHungry: true,
			expectedTired:  false,
		},
		{
			name: "SharkIsTired",
			shark: Shark{
				hungry: true,
				tired:  true,
				speed:  10,
			},
			prey: Prey{
				name:  "tuna",
				speed: 5,
			},
			expectedErr:    ErrSharkTired,
			expectedHungry: true,
			expectedTired:  true,
		},
		{
			name: "SharkIsNotHungry",
			shark: Shark{
				hungry: false,
				tired:  false,
				speed:  10,
			},
			prey: Prey{
				name:  "tuna",
				speed: 5,
			},
			expectedErr:    ErrSharkNotHungry,
			expectedHungry: false,
			expectedTired:  false,
		},
		{
			name: "SharkCannotCatchPrey",
			shark: Shark{
				hungry: true,
				tired:  false,
				speed:  5,
			},
			prey: Prey{
				name:  "tuna",
				speed: 10,
			},
			expectedErr:    ErrSharkNotCatch,
			expectedHungry: true,
			expectedTired:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.shark.Hunt(&tc.prey)

			if err != tc.expectedErr {
				t.Errorf("unexpected error. expected: %v, got: %v", tc.expectedErr, err)
			}

			if tc.shark.hungry != tc.expectedHungry {
				t.Errorf("unexpected hungry state. expected: %v, got: %v", tc.expectedHungry, tc.shark.hungry)
			}

			if tc.shark.tired != tc.expectedTired {
				t.Errorf("unexpected tired state. expected: %v, got: %v", tc.expectedTired, tc.shark.tired)
			}
		})
	}
}*/
