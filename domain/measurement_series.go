package domain

import "time"

type MeasurementSeries struct {
	ID                uint64     `sql:"AUTO_INCREMENT"; json:"id"`
	StartTimestamp    time.Time  `sql:"NOT NULL"; json:"id"`
	EndTimestamp      *time.Time `sql:"size:29"; json:"endTimestamp"`
	MeasurementEntity []Measurement
}
