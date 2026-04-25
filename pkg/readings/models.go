package readings

type Reading struct {
	SensorId string  `json:"sensorId"`
	Moisture float32 `json:"moisture"`
}
