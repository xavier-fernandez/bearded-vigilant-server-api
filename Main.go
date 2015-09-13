package main

import (
	"github.com/environ-vigilant/persistance"
	"github.com/environ-vigilant/sensor"
	"github.com/environ-vigilant/sensor/humidity"
	"time"
)

const (
	APP_NAME    = "enviro-vigilant"
	APP_VERSION = "0.0.1"
)

func main() {
	println("Loading", APP_NAME, "version", APP_VERSION)
	persistance.InitDatabase()
	grovePi := sensor.InitGrovePi(0x04)
	dht := humidity.InitDHT(grovePi, sensor.D3)
	for true {
		time.Sleep(600 * time.Millisecond)
		dht.ReadAndStoreSensorData()
	}
}
