package readings

type ReadingStore interface {
	Add(reading AddReading) error
	List() ([]Reading, error)
	Latest() ([]Reading, error)
	GetSensor(sensorId string) (Sensor, error)
	SetSensorName(sensorId string, sensorName string) error
}
