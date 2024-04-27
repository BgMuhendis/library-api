package database

import (
	"os"
	"testing"

)


func TestConnectDB(t *testing.T){
		err:=os.Setenv("DATABASE_URL","host=127.0.0.1 user=postgres password=postgres dbname=library sslmode=disable")

		if err !=nil {
			t.Errorf("Set edilemedi")
		}
		db,err := ConnectDB()

		if err != nil {
			t.Error("DB is not connection")
		}

		t.Log("DB ye bağlanıldı",db.Db.Config.ConnPool)

}