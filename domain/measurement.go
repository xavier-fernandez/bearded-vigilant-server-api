package domain

import "time"

type Measurement struct {
	ID             uint64     `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	Value          float64    `sql:"NOT NULL";`
	BinSize        uint16     `sql:"NOT NULL";`
	StartTimestamp time.Time  `sql:"NOT NULL"; json:"id"`
	EndTimestamp   *time.Time `json:"endTimestamp"`
	Latitude       float32    `sql:"size:4"; json:"latitude"`
	Longitude      float32    `sql:"size:4"; json:"longitude"`
}
