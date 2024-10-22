package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() {

	// Guanakan Ketika Deploy
	dsn := "root:edTyDluWiIGDqijgpDygEVvoqdPNXfzL@tcp(autorack.proxy.rlwy.net:24422)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := os.Getenv("DB_USER") + ":" +
	// 	os.Getenv("DB_PASS") + "@(" +
	// 	os.Getenv("DB_HOST") + ":" +
	// 	os.Getenv("DB_PORT") + ")/" +
	// 	os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// db.Migrate(db.DB)
	DBConn = db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Posting{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Like{})
}
