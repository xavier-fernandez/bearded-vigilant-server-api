package main

import (
	"github.com/environ-vigilant/persistance"
)

const APP_NAME = "enviro-vigilant"
const APP_VERSION = "0.0.1"

func main() {
	println("Loading", APP_NAME, "version", APP_VERSION)
	persistance.InitDatabase()
}
