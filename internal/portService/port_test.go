package portService

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPortDeserialisation(t *testing.T) {
	type testCase struct {
		name         string
		input        string
		expectedPort Port
	}

	// Here I would continue with a few examples from the data source. I would also
	// check using invalid source data to make sure that the desrialisation fails.
	testCases := []testCase{
		{
			name: "valid port",
			input: `{
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }`,
			expectedPort: Port{
				Name:    "Ajman",
				City:    "Ajman",
				Country: "United Arab Emirates",
				Coordinates: Coordinate{
					Latitude:  55.5136433,
					Longitude: 25.4052165,
				},
				Province: "Ajman",
				Unlocs:   []string{"AEAJM"},
				Code:     "52000",
				Timezone: "Asia/Dubai",
				Alias:    []string{},
				Regions:  []string{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var gotPort Port
			require.NoError(t, json.Unmarshal([]byte(tc.input), &gotPort))
			assert.Equal(t, tc.expectedPort, gotPort)
		})
	}
}
