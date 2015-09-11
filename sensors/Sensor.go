package sensors

type Sensor interface {
	ReadAndStoreSensorData() float64
}
