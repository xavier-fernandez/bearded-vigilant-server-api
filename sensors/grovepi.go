package sensors

import (
	"fmt"
	"github.com/mrmorphic/hwio"
	"time"
)

const (
	A0 = 0  //GrovePi Pin - Analog 0
	A1 = 1  //GrovePi Pin - Analog 1
	A2 = 2  //GrovePi Pin - Analog 2
	D2 = 2  //GrovePi Pin - Digital 2
	D3 = 3  //GrovePi Pin - Digital 3
	D4 = 4  //GrovePi Pin - Digital 4
	D5 = 5  //GrovePi Pin - Digital 5
	D6 = 6  //GrovePi Pin - Digital 6
	D7 = 7  //GrovePi Pin - Digital 7
	D8 = 8  //GrovePi Pin - Digital 8

	DIGITAL_READ = 1
	DIGITAL_WRITE = 2
	ANALOG_READ = 3
	ANALOG_WRITE = 4
	PIN_MODE = 5
)

// This struct provides the basic functions for using the GrovePi
type GrovePi struct {
	i2cModule hwio.I2CModule
	i2cDevice hwio.I2CDevice
}

// Establish the connection with a GrovePi.
// This method should be called for initializing the GrovePi struct.
// It returns a reference to the GrovePi struct.
func InitGrovePi(address int) *GrovePi {
	grovePi := new(GrovePi)
	m, err := hwio.GetModule("i2c")
	if err != nil {
		fmt.Printf("could not get i2c module: %s\n", err)
		return nil
	}
	grovePi.i2cModule = m.(hwio.I2CModule)
	grovePi.i2cModule.Enable()

	grovePi.i2cDevice = grovePi.i2cModule.GetDevice(address)
	return grovePi
}

// Close the GrovePi connection
func (grovePi *GrovePi) CloseDevice() {
	grovePi.i2cModule.Disable()
}

// Does an analog read on the selected pin.
// We can use 'grovepi.go' constants, we can access the
// pins easily calling 'grovepi.D1', for example.
// It will return the sensor value if it was possible to parse the value.
// It will return an error if some problem happened when reading the value.
func (grovePi *GrovePi) analogRead(pin byte) (int, error) {
	b := []byte{ANALOG_READ, pin, 0, 0}
	err := grovePi.i2cDevice.Write(1, b)
	if err != nil {
		return 0, err
	}
	time.Sleep(100 * time.Millisecond)
	grovePi.i2cDevice.ReadByte(1)
	val, err2 := grovePi.i2cDevice.Read(1, 4)
	if err2 != nil {
		return 0, err
	}
	var v1 int = int(val[1])
	var v2 int = int(val[2])
	return ((v1 * 256) + v2), nil
}

// Does an digital read on the selected pin.
// We can use 'grovepi.go' constants, we can access the
// pins easily calling 'grovepi.D1', for example.
// It will return the sensor value if it was possible to parse the value.
// It will return an error if some problem happened when reading the value.
func (grovePi *GrovePi) digitalRead(pin byte) (byte, error) {
	b := []byte{DIGITAL_READ, pin, 0, 0}
	err := grovePi.i2cDevice.Write(1, b)
	if err != nil {
		return 0, err
	}
	time.Sleep(100 * time.Millisecond)
	val, err2 := grovePi.i2cDevice.ReadByte(1)
	if err2 != nil {
		return 0, err2
	}
	return val, nil
}

// Does a digital write on the selected pin.
// We can use 'grovepi.go' constants, we can access the
// pins easily calling 'grovepi.D1', for example.
// It will return an error if some problem happened when writing the value.
func (grovePi *GrovePi) digitalWrite(pin byte, val byte) error {
	b := []byte{DIGITAL_WRITE, pin, val, 0}
	err := grovePi.i2cDevice.Write(1, b)
	time.Sleep(100 * time.Millisecond)
	if err != nil {
		return err
	}
	return nil
}

// Modifies the mode of a GrovePi pin
// If mode is string is equal to "output" it will set the mode to "output".
// Every other mode will set the mode to "input"
// It will return an error if some problem happened when changing the pin mode.
func (grovePi *GrovePi) pinMode(pin byte, mode string) error {
	var b []byte
	if mode == "output" {
		b = []byte{PIN_MODE, pin, 1, 0}
	} else {
		b = []byte{PIN_MODE, pin, 0, 0}
	}
	err := grovePi.i2cDevice.Write(1, b)
	time.Sleep(100 * time.Millisecond)
	if err != nil {
		return err
	}
	return nil
}
