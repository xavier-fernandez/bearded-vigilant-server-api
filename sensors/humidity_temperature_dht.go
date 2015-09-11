package sensors

import (
	"time"
	"unsafe"
)

const (
	DHT_READ = 40
)

// Reads the temperature and humidity information from the selected GrovePi pin.
// It returns the temperature and the humidity (in this order) as floats.
// It returns an error in case some error occurs when extracting the data.
func (grovePi *GrovePi) ReadDHT(pin byte) (float32, float32, error) {
	b := []byte{DHT_READ, pin, 1, 0}
	rawData, err := grovePi.readDHTRawData(b)
	if err != nil {
		return 0, 0, err
	}
	temperatureData := rawData[1:5]

	temperatureInt := int32(temperatureData[0]) | int32(temperatureData[1]) << 8 | int32(temperatureData[2]) << 16 | int32(temperatureData[3]) << 24
	temperature := (*(*float32)(unsafe.Pointer(&temperatureInt)))

	humidityData := rawData[5:9]
	humidityInt := int32(humidityData[0]) | int32(humidityData[1]) << 8 | int32(humidityData[2]) << 16 | int32(humidityData[3]) << 24
	humidity := (*(*float32)(unsafe.Pointer(&humidityInt)))
	return temperature, humidity, nil
}

// Extracts the current rawData from the connected DHT sensor
func (grovePi *GrovePi) readDHTRawData(cmd []byte) ([]byte, error) {
	err := grovePi.i2cDevice.Write(1, cmd)
	if err != nil {
		return nil, err
	}
	time.Sleep(600 * time.Millisecond)
	grovePi.i2cDevice.ReadByte(1)
	time.Sleep(100 * time.Millisecond)
	raw, err := grovePi.i2cDevice.Read(1, 9)
	if err != nil {
		return nil, err
	}
	return raw, nil
}
