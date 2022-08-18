package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type dbConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() error {
	envLoadErr := godotenv.Load("../config/.env")
	if envLoadErr != nil {
		log.Fatal("Could not set the environment")
		return envLoadErr
	}

	// connection string object
	conn := dbConnection{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	// create connection string
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", conn.Host, conn.Port, conn.User, conn.Password, conn.DBName)

	var err error
	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Could not open database connection")
		return err
	}

	// test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not establish database connection")
		return err
	}

	return nil

}
