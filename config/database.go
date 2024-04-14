package config

import (
	_"fmt"	
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

/* 	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"),os.Getenv("POSTGRES_PORT"),os.Getenv("POSTGRES_USER"),os.Getenv("POSTGRES_PASSWORD"),os.Getenv("POSTGRES_DB"))
 */
 dbInfo := os.Getenv("DATABASE_URL")
 db,err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}


	return db
}
