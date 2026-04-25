package readings

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type ReadingDbStore struct {
	db *sql.DB
}

func NewReadingDbStore(dbPath string) *ReadingDbStore {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Println("Unable to open sqlite db")
	}

	return &ReadingDbStore{
		db,
	}
}

func (s *ReadingDbStore) Add(reading Reading) error {
	statement, err := s.db.Prepare("insert into readings(sensorId, moisture) values (?, ?)")
	defer statement.Close()
	if err != nil {
		return err
	}

	_, err = statement.Exec(reading.SensorId, reading.Moisture)

	if err != nil {
		return err
	}

	return nil
}

func (s *ReadingDbStore) List() ([]Reading, error) {
	var readings []Reading

	rows, err := s.db.Query("select sensorId, moisture from readings")
	if err != nil {
		return readings, err
	}

	for rows.Next() {
		var r Reading
		err := rows.Scan(&r.SensorId, &r.Moisture)
		if err != nil {
			return readings, err
		}
		readings = append(readings, r)
	}

	return readings, nil
}
