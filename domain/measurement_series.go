package domain

import "time"

type MeasurementSeries struct {
	ID                uint64     `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	StartTimestamp    time.Time  `sql:"NOT NULL"; json:"id"`
	EndTimestamp      *time.Time `json:"endTimestamp"`
	MeasurementEntity []Measurement
}
