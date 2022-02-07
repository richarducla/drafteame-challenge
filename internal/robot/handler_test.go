package robot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobot_Process(t *testing.T) {
	testCases := []struct {
		name             string
		instructions     []string
		expectedResponse []Robot
		expectedError    bool
		messageError     string
	}{
		{
			name:         "successful processing",
			instructions: []string{"5 3", "1 1 E", "RFRFRFRF", "3 2 N", "FRRFLLFFRRFLL", "0 3 W", "LLFFFLFLFL"},
			expectedResponse: []Robot{
				{
					X:           1,
					Y:           1,
					Orientation: "E",
				},
				{
					X:           3,
					Y:           3,
					Orientation: "N",
					IsLost:      "LOST",
				},
				{
					X:           2,
					Y:           3,
					Orientation: "S",
				},
			},
		},
		{
			name:             "grid invalid",
			instructions:     []string{"89 3", "1 1 E", "RFRFRFRF"},
			expectedResponse: []Robot{},
			expectedError:    true,
			messageError:     "the grid is invalid",
		},
		{
			name:             "invalid coordinate for x-axis in grid",
			instructions:     []string{"asd 3", "1 1 E", "RFRFRFRF"},
			expectedResponse: []Robot{},
			expectedError:    true,
			messageError:     "coordinate grid x invalid",
		},
		{
			name:             "invalid coordinate for y-axis in grid",
			instructions:     []string{"5 asd", "1 1 E", "RFRFRFRF"},
			expectedResponse: []Robot{},
			expectedError:    true,
			messageError:     "coordinate grid y invalid",
		},
		{
			name:             "invalid coordinate for x-axis in robot",
			instructions:     []string{"5 3", "asd 1 E", "RFRFRFRF"},
			expectedResponse: []Robot{},
			expectedError:    true,
			messageError:     "coordinate x invalid",
		},
		{
			name:             "invalid coordinate for y-axis in robot",
			instructions:     []string{"5 3", "1 asd E", "RFRFRFRF"},
			expectedResponse: []Robot{},
			expectedError:    true,
			messageError:     "coordinate y invalid",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := Process(testCase.instructions)
			if testCase.expectedError {
				assert.EqualError(t, err, testCase.messageError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedResponse, result)
			}
		})
	}
}
