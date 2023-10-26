package services

import (
	"errors"

	"github.com/devminnu/tek-job-portal/internal/database"
)

type Service interface {
	AutoMigrate() error
}

type service struct {
	Service
	db *database.Conn
}

func New(db *database.Conn) (Service, error) {
	// We check if the database instance is nil, which would indicate an issue.
	if db == nil {
		return nil, errors.New("please provide a valid db connection")
	}

	return &service{db: db}, nil
}
