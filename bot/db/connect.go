package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() {

	var err error
	Psql, err = gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION_STRINGS")))

	if err != nil {
		panic(err)
	} else {
		log.Print("[INFO] DB connected!")
	}

	Psql.Logger = Psql.Logger.LogMode(logger.Info)
	return

}
