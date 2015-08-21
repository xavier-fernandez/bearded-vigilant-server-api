package domain

import "time"

type MeasurementSeries struct {
	ID                uint32     `sql:"AUTO_INCREMENT"; json:"id"`
	StartTimestamp    time.Time  `sql:"size:29"; json:"id"`
	EndTimestamp      *time.Time `sql:"size:29"; json:"endTimestamp"`
	MeasurementEntity []Measurement
}
