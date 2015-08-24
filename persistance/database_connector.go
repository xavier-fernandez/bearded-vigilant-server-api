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
	// Initialize the database connection
	database.DB()
	// Enable database Logging.
	database.LogMode(true)
	// Creates the database tables.
	database.CreateTable(&domain.MeasurementSeries{})
	database.CreateTable(&domain.Measurement{})
	database.CreateTable(&domain.Sensor{})
	database.CreateTable(&domain.SensorType{})
	db = database
	initForeignKeys()
	return nil
}

func initForeignKeys() {
	// Initializes the measurements table foreign keys
	measurementModel := db.Model(&domain.Measurement{})
	measurementModel.AddForeignKey("measurement_series_id", "measurement_series(id)", "CASCADE", "CASCADE")
	// Initializes the measurements_series table foreign key
	measurementSeriesModel := db.Model(&domain.MeasurementSeries{})
	measurementSeriesModel.AddForeignKey("sensor_id", "sensors(id)", "CASCADE", "CASCADE")
	measurementSeriesModel.AddForeignKey("sensor_type_id", "sensor_types(id)", "CASCADE", "CASCADE")
}
