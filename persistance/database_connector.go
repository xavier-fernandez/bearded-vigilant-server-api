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

// Obtains a database connector, if available.
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
	database.CreateTable(&domain.Sensor{})
	database.CreateTable(&domain.SensorType{})
	db = database
	initMeasurementsTable()
	return nil
}

func initMeasurementsTable() {
	measurementEntity := &domain.Measurement{}
	db.CreateTable(measurementEntity)
	db.Model(measurementEntity).AddForeignKey("measurement_series_id", "measurement_series(id)", "CASCADE", "CASCADE")
}
