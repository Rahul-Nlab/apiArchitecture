package database

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

const (
	dbHost     = "192.168.29.174"
	dbPort     = "5434"
	dbUser     = "rahul"
	dbPassword = "rahul"
	dbName     = "rahul"
)

func Setup() *sqlx.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sqlx.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("There was an error while connecting to the database.")
		return nil
	}
	return db
}
