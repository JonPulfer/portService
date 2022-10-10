package portService

import (
	"encoding/json"
	"fmt"
)

type Port struct {
	Name        string     `json:"name"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
	Timezone    string     `json:"timezone"`
	Unlocs      []string   `json:"unlocs"`
	Alias       []string   `json:"alias"`
	Regions     []string   `json:"regions"`
	Province    string     `json:"province"`
	Coordinates Coordinate `json:"coordinates"`
	Code        string     `json:"code"`
}

type Coordinate struct {
	Latitude  float32
	Longitude float32
}

func (c *Coordinate) UnmarshalJSON(bytes []byte) error {
	var got []float32
	err := json.Unmarshal(bytes, &got)
	if err != nil {
		return fmt.Errorf("failed to deserialise coordinates: %w", err)
	}
	c.Latitude = got[0]
	c.Longitude = got[1]
	return nil
}

func NewCoordinate(x, y float32) *Coordinate {
	return &Coordinate{Latitude: x, Longitude: y}
}
