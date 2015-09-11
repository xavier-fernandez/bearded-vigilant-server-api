package humidity

import (
	"fmt"
	"github.com/environ-vigilant/sensors"
	"time"
	"unsafe"
)

const (
	DHT_READ = 40
)

// This struct provides the basic functions for using the GrovePi
type DHT struct {
	grovePi *sensors.GrovePi
	pin     byte
}

func InitDHT(grovePi *sensors.GrovePi, pinNumber byte) *DHT {
	sensor := new(DHT)
	sensor.grovePi = grovePi
	sensor.pin = pinNumber
	return sensor
}

// Obtains and stores humidity and temperature data
func (sensor *DHT) ReadAndStoreSensorData() {
	temperature, humidity, err := sensor.readDHT(sensor.pin)
	sensorPin := sensor.pin
	if err != nil {
		fmt.Errorf("Following error was thrown when obtaining the RHT data from the sensor ", sensorPin, " -> ", err)
	} else {
		grovePiAddress := sensor.grovePi.Address()
		println("GrovePi: ", grovePiAddress, "Pin: ", sensorPin, " Temperature: ", temperature, " Humidity: ", humidity)
		//TODO: Store data into the database
	}
}

// Reads the temperature and humidity information from the selected GrovePi pin.
// It returns the temperature and the humidity (in this order) as floats.
// It returns an error in case some error occurs when extracting the data.
func (sensor *DHT) readDHT(pin byte) (float32, float32, error) {
	b := []byte{DHT_READ, pin, 1, 0}
	rawData, err := sensor.readDHTRawData(b)
	if err != nil {
		return 0, 0, err
	}
	temperatureData := rawData[1:5]

	temperatureInt := int32(temperatureData[0]) | int32(temperatureData[1])<<8 | int32(temperatureData[2])<<16 | int32(temperatureData[3])<<24
	temperature := (*(*float32)(unsafe.Pointer(&temperatureInt)))

	humidityData := rawData[5:9]
	humidityInt := int32(humidityData[0]) | int32(humidityData[1])<<8 | int32(humidityData[2])<<16 | int32(humidityData[3])<<24
	humidity := (*(*float32)(unsafe.Pointer(&humidityInt)))
	return temperature, humidity, nil
}

// Extracts the current rawData from the connected DHT sensor
func (sensor *DHT) readDHTRawData(cmd []byte) ([]byte, error) {
	err := sensor.grovePi.I2CDevice().Write(1, cmd)
	if err != nil {
		return nil, err
	}
	time.Sleep(600 * time.Millisecond)
	sensor.grovePi.I2CDevice().ReadByte(1)
	time.Sleep(100 * time.Millisecond)
	raw, err := sensor.grovePi.I2CDevice().Read(1, 9)
	if err != nil {
		return nil, err
	}
	return raw, nil
}
