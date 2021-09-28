package infrastructure

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	Table      *gorm.DB
	Connection *sql.DB
}

var PGdatabase = Adapter{}

func (e *Environment) InitDatabase() error {
	db, err := gorm.Open(postgres.Open(e.Databases["db"].Connection), &gorm.Config{})
	if err != nil {
		return err
	}
	connection, err := db.DB()
	if err != nil {
		return err
	}
	PGdatabase = Adapter{Table: db, Connection: connection}
	log.Println("Database load successfully!")
	return nil
}
