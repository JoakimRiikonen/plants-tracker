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

func (s *ReadingDbStore) Add(reading AddReading) error {
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

	query := `
	SELECT r.sensorId, s.sensorName, r.moisture, r.timestamp
	FROM readings r
	LEFT JOIN sensors s ON r.sensorId = s.sensorId`

	rows, err := s.db.Query(query)
	if err != nil {
		return readings, err
	}

	for rows.Next() {
		var r Reading
		err := rows.Scan(&r.SensorId, &r.SensorName, &r.Moisture, &r.Timestamp)
		if err != nil {
			return readings, err
		}
		readings = append(readings, r)
	}

	return readings, nil
}

func (s *ReadingDbStore) Latest() ([]Reading, error) {
	var readings []Reading

	query := `
	SELECT r1.sensorId, s.sensorName, r1.moisture, MAX(r1.timestamp)
	FROM readings r1
	LEFT JOIN sensors s ON r1.sensorId = s.sensorId
	GROUP BY r1.sensorId
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return readings, err
	}

	for rows.Next() {
		var r Reading
		err := rows.Scan(&r.SensorId, &r.SensorName, &r.Moisture, &r.Timestamp)
		if err != nil {
			return readings, err
		}
		readings = append(readings, r)
	}

	return readings, nil
}
