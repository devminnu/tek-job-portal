package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Conn is our main struct, including the database instance for working with data.
type Conn struct {
	// db is an instance of the SQLite database.
	db *gorm.DB
}

// NewService is the constructor for the Conn struct.
func New(db *gorm.DB) (*Conn, error) {

	// We check if the database instance is nil, which would indicate an issue.
	if db == nil {
		return nil, errors.New("please provide a valid connection")
	}

	// We initialize our service with the passed database instance.
	s := &Conn{db: db}

	return s, nil
}

func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=diwakar password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
