package readings

type SensorDbModel struct {
	SensorId   string
	SensorName float32
}

type ReadingDbModel struct {
	SensorId string
	Moisture float32
}
