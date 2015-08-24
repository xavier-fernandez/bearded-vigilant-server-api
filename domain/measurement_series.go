package domain

import "time"

type MeasurementSeries struct {
	ID                uint64         `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	StartTimestamp    time.Time      `sql:"NOT NULL"; json:"startTimestamp"`
	EndTimestamp      *time.Time     `json:"endTimestamp"`
	SensorId          uint64         `gorm:"column:sensor_id"`
	Sensor            Sensor
	MeasurementEntity []Measurement
	SensorTypeId 	  uint64         `gorm:"column:sensor_type_id"`
	SensorType		  SensorType
}