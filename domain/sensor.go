package domain

type Sensor struct {
	ID         uint64 `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	IsI2c      bool   `sql:"NOT NULL"; json:"isI2C"`
	PinNumber  uint32 `sql:"NOT NULL"; json:"pinNumber"`
	GrovePiRow uint8  `sql:"NOT NULL"; json:"grovePiRow"`
}
