package readings

type Reading struct {
	SensorId string  `json:"sensorId"`
	Moisture float64 `json:"moisture"`
}
