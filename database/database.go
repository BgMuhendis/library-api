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

func ConnectDB() (*DBInfo,error) {
 dbInfo := os.Getenv("DATABASE_URL")
 db,err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})

	if err != nil {
		return nil,err
	}


	return &DBInfo {
		Db:db,
	},nil
}

func (d *DBInfo) DBClose()  {
		 connection, _ := d.Db.DB()
		 connection.Close() 
}

func (d *DBInfo) Runmigrate(tableName string)  {
		d.Db.Table(tableName).AutoMigrate(&entity.Book{})
}
