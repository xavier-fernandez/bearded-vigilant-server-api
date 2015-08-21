package domain

type Sensor struct {
	ID         uint64 `sql:"AUTO_INCREMENT"; sql:"NOT NULL"; json:"id"`
	IsAnalog   bool   `sql:"NOT NULL"; json:"isAnalog"`
	IsDigital  bool   `sql:"NOT NULL"; json:"isDigital"`
	IsPwm      bool   `sql:"NOT NULL"; json:"isPAM"`
	IsI2c      bool   `sql:"NOT NULL"; json:"isI2C"`
	PinNumber  uint32 `sql:"NOT NULL"; json:"pinNumber"`
	GrovePiRow uint8  `sql:"NOT NULL"; json:"grovePiRow"`
}
