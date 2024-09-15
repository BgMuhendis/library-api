package database

import (
	"library/internal/app/model/entity"
	"os"
	"fmt"

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
		 connection, err := d.Db.DB()

		 if err!=nil {
			fmt.Println(nil)
		 }
		 connection.Close() 
}

func (d *DBInfo) RunMigrate(tableName string)  {
		d.Db.Table(tableName).AutoMigrate(&entity.Book{})
}
