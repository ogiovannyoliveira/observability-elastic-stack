package database

import (
	"log"

	"github.com/giovannylucas/observability-elastic-stack/golang-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := config.GetDSN()

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Error when trying to connect on database", err.Error())
	}

	log.Println("Connected on database")
}