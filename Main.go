package main

import (
	"github.com/environ-vigilant/persistance"
	"github.com/environ-vigilant/sensors"
	"time"
	"github.com/environ-vigilant/sensors/humidity"
)

const APP_NAME = "enviro-vigilant"
const APP_VERSION = "0.0.1"

func main() {
	println("Loading", APP_NAME, "version", APP_VERSION)
	persistance.InitDatabase()
	grovePi := sensors.InitGrovePi(0x04)
	dht := humidity.InitDHT(grovePi, sensors.D3)
	for true {
		time.Sleep(600 * time.Millisecond)
		dht.ReadAndStoreSensorData()
	}
}
