package readings

type SensorDbModel struct {
	SensorId   string
	SensorName float64
}

type ReadingDbModel struct {
	SensorId string
	Moisture float64
}
