package config

import (
	"fmt"
	"os"
)

var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var DB_HOST string
var DB_PORT string

func init() {
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
}

func GetDSN() string {
	return "host=127.0.0.1 user=developer password=dv1010aa dbname=classes_scheduler port=5432 sslmode=disable"
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)
}