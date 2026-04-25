package readings

type ReadingStore interface {
	Add(reading Reading) error
	List() ([]Reading, error)
}
