package sensor

type SensorFactory struct {
	Sensors map[string]string
}

func InitSensorFactory() *SensorFactory {
	factory := new(SensorFactory)
	factory.Sensors = make(map[string]string)
	return factory
}
