package domain

import "time"

type Measurement struct {
	ID             uint32     `sql:"AUTO_INCREMENT"; json:"id"`
	StartTimestamp time.Time  `sql:"size:29"; json:"id"`
	EndTimestamp   *time.Time `sql:"size:29"; json:"endTimestamp"`
	Latitude       float32    `sql:"size:4"; json:"latitude"`
	Longitude      float32    `sql:"size:4"; json:"longitude"`
}
