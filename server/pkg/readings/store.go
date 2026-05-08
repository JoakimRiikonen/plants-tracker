package readings

type ReadingStore interface {
	Add(reading AddReading) error
	List() ([]Reading, error)
	Latest() ([]Reading, error)
}
