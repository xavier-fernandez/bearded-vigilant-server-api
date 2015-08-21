package domain

type SensorUnit struct {
	ID   uint64 `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	Name string `sql:"UNIQUE; NOT NULL"; json:"name"`
}
