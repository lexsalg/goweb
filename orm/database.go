package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/lexsalg/goweb/config"
)

var db *gorm.DB

func CreateConnection() {
	dsn := config.DsnDB()
	if conn, err := gorm.Open("mysql", dsn); err != nil {
		panic(err)
	} else {
		db = conn

	}
}
func CloseConnection() {
	_ = db.Close()
}

func CreateTables() {
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})
}