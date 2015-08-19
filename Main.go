package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const APP_NAME = "enviro-vigilant"
const APP_VERSION = "0.0.1"

const DATABASE_LOCATION = "databases/enviro-gatt.db"

func main() {
	println("Loading", APP_NAME, "version", APP_VERSION)
	db, err := gorm.Open("sqlite3", DATABASE_LOCATION)
	if err != nil {
		fmt.Errorf("The following error was thrown when initializing the database -> ", err)
	}
	// Initialize the database connection
	db.DB()
	// Enable database Logging.
	db.LogMode(true)
}
