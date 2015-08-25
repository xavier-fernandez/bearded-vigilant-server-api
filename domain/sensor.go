package domain

type Sensor struct {
	ID          uint64       `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	PinNumber   uint32       `sql:"NOT NULL"; json:"pinNumber"`
	GrovePiRow  uint8        `sql:"NOT NULL"; json:"grovePiRow"`
	SensorTypes []SensorType `gorm:"many2many:sensor_type_sensor;"`
}
