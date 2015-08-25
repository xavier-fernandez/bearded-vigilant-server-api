package domain

type SensorOutputType struct {
	ID   uint64 `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	Name string `sql:"UNIQUE; NOT NULL"; json:"name"`
	Unit string `sql:"UNIQUE; NOT NULL"; json:"unit"`
}
