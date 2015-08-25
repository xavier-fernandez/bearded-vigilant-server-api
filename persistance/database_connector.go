package persistance

import (
	"fmt"
	"github.com/environ-vigilant/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const databaseUser = "envirodb"
const databasePassword = "H8L@r3@{~?&~}w|J33$6!837?l"

var db gorm.DB

// Initializes the database, throws an error if an error is produced
func InitDatabase() error {
	database, err := gorm.Open("postgres", "user=Xavi dbname=Xavi sslmode=disable")
	if err != nil {
		fmt.Errorf("The following error was thrown when initializing the database -> ", err)
		return err
	}
	db = database
	initDatabaseTables()
	db.DB()
	db.LogMode(true)
	initForeignKeys()
	return nil
}

func initDatabaseTables() {
	db.CreateTable(&domain.Measurement{})
	db.CreateTable(&domain.Sensor{})
	db.CreateTable(&domain.SensorType{})
	db.CreateTable(&domain.SensorOutputType{})
}

func initForeignKeys() {
	// Initializes the measurements table foreign keys
	measurementModel := db.Model(&domain.Measurement{})
	measurementModel.AddForeignKey("sensor_id", "sensors(id)", "CASCADE", "CASCADE")
	measurementModel.AddForeignKey("sensor_output_type_id", "sensor_output_types(id)", "CASCADE", "CASCADE")
	// Initializes the sensor_type table foreign key
	sensorTypeModel := db.Model(&domain.SensorType{})
	sensorTypeModel.AddForeignKey("sensor_output_type_id", "sensor_output_types(id)", "CASCADE", "CASCADE")
}
