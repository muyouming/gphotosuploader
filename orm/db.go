package orm

import (
	"sync"

	"github.com/garfieldkwong/gphotosuploader/orm/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// The DB type
type DB struct {
	Connection *gorm.DB
}

var instance *DB
var once sync.Once

// GetInstance Retrieve the singleton instance of DB
func GetInstance() *DB {
	once.Do(func() {
		instance = &DB{}
		var err error
		instance.Connection, err = gorm.Open("sqlite3", "./db.sqlite3?loc=UTC")
		if err != nil {
			panic("failed to connect database")
		}
		instance.Connection.AutoMigrate(&models.File{})
		if err != nil {
			panic("failed to connect database")
		}
	})

	return instance
}
