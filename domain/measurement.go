package domain

import "time"

type Measurement struct {
	ID                 uint64 `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	SensorOutputTypeId uint64 `sql:"NOT NULL";`
	SensorOutputType   SensorOutputType
	SensorId           uint64 `sql:"NOT NULL";`
	Sensor             Sensor
	Value              float64    `sql:"NOT NULL"; json:"value"`
	BinSize            uint16     `sql:"NOT NULL"; json:"binSize"`
	StartTimestamp     time.Time  `sql:"NOT NULL"; json:"startTimestamp"`
	EndTimestamp       *time.Time `sql:"NOT NULL"; json:"endTimestamp"`
	Latitude           float32    `json:"latitude"`
	Longitude          float32    `json:"longitude"`
}
