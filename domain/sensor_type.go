package domain

type SensorType struct {
	ID                 uint64 `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	Name               string `sql:"UNIQUE; NOT NULL"; json:"name"`
	IsAnalog           bool   `sql:"NOT NULL"; json:"isAnalog"`
	IsDigital          bool   `sql:"NOT NULL"; json:"isDigital"`
	IsPwm              bool   `sql:"NOT NULL"; json:"isPAM"`
	IsI2c              bool   `sql:"NOT NULL"; json:"isI2C"`
	SensorOutputTypeId uint64 `sql:"NOT NULL";`
	SensorOutputType   SensorOutputType
}
