package sensor

// Represents the abstraction of a Sensor that all the project sensors classes must implement.
type Sensor interface {
	// Tries to read the current data from the sensor, storing the result in the database.
	ReadAndStoreSensorData() error
	// Gets the sensor name.
	Name() string
}
