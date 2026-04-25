package readings

import "fmt"

type ReadingMemStore struct {
	list []Reading
}

func NewReadingMemStore() *ReadingMemStore {
	list := make([]Reading, 0)
	return &ReadingMemStore{
		list,
	}
}

func (s *ReadingMemStore) Add(reading Reading) error {
	fmt.Println("Adding to list...")
	s.list = append(s.list, reading)
	return nil
}

func (s *ReadingMemStore) List() ([]Reading, error) {
	return s.list, nil
}
