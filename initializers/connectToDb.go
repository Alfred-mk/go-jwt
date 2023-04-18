package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "go_jwt:cxXr33#fEDNGvt@tcp(db4free.net:3306)/go_jwt?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}
}
