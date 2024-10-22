package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() {
	// Hardcoded DSN for MySQL connection
	dsn := "root:edTyDluWiIGDqijgpDygEVvoqdPNXfzL@tcp(junction.proxy.rlwy.net:51729)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Perform migrations
	Migrate(db)

	DBConn = db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Posting{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Like{})
}
