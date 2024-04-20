package database

import (
	"library/model/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type DBInfo struct {
		Db  *gorm.DB
}

func ConnectDB() *DBInfo {
 dbInfo := os.Getenv("DATABASE_URL")
 db,err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}


	return &DBInfo {
		Db:db,
	}
}

func (d *DBInfo) DBClose()  {
		 connection, _ := d.Db.DB()
		 connection.Close() 
}

func (d *DBInfo) Automigrate(tableName string)  {
		d.Db.Table("books").AutoMigrate(&entity.Book{})
}
