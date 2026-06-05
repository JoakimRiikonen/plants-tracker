package readings

type Reading struct {
	SensorId   string
	SensorName string
	Moisture   float64
	Timestamp  string
}

type Sensor struct {
	SensorId   string
	SensorName string
}

type AddReading struct {
	SensorId string
	Moisture float64
}
