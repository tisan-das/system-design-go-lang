package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection interface {
	connect(hostname, user, password, dbname string, port int)
}

type postgresConnection struct{}

func NewPostgresDBConnection(hostname, user, password, dbname string, port int) (*gorm.DB, error) {
	var postgresDB postgresConnection
	return postgresDB.connect(hostname, user, password, dbname, port)
}

func (db *postgresConnection) connect(hostname, user, password, dbname string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		hostname, user, password, dbname, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
