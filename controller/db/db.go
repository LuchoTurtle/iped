package db

import (
	"fmt"
	"os"

	"github.com/ipp-ementa/iped/model/canteen"
	"github.com/ipp-ementa/iped/model/dish"
	"github.com/ipp-ementa/iped/model/menu"
	"github.com/ipp-ementa/iped/model/school"

	"github.com/jinzhu/gorm"
	// Requires to import sqlite dialect package to use and open sqlite3 database
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Db holds the state of the database connection
var Db *gorm.DB

// Initializes database once go program starts or package is imported
// Panics if database couldn't be open
func init() {

	conn := os.Getenv("IPEW_CONNECTION_STRING")

	odb, err := gorm.Open("sqlite3", conn)

	if err != nil {
		panic(fmt.Sprintf("Database '%s' couldn't be open due to: %s", conn, err))
	}

	odb.Exec("PRAGMA foreign_keys = ON")

	odb.AutoMigrate(&school.School{})

	odb.AutoMigrate(&canteen.Canteen{})

	odb.AutoMigrate(&canteen.MenuEntry{})

	odb.AutoMigrate(&menu.Menu{})

	odb.AutoMigrate(&dish.Dish{})

	Db = odb
}
