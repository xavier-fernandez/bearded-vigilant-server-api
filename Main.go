package main

import (
	"github.com/environ-vigilant/persistance"
	"github.com/environ-vigilant/sensors"
	"time"
)

const APP_NAME = "enviro-vigilant"
const APP_VERSION = "0.0.1"

func main() {
	println("Loading", APP_NAME, "version", APP_VERSION)
	persistance.InitDatabase()
	grovePi := sensors.InitGrovePi(0x04)
	for true {
		time.Sleep(600 * time.Millisecond)
		temperature, humidity, err := grovePi.ReadDHT(sensors.D3)
		if err != nil {
			println("The following error was thrown:", err)
		} else {
			println("Temperature: ", temperature, " Humidity: ", humidity)
		}
	}
}
