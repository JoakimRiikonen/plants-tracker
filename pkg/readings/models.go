package readings

type Reading struct {
	SensorId  string
	Moisture  float64
	Timestamp string
}

type AddReading struct {
	SensorId string
	Moisture float64
}
