package models

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/provalidator-nodereal-api-marketplace/log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var DbStr string

func ConnectDatabase() {
	var err error

	DbStr = os.Getenv("MYSQL_USER_ID") + ":" + os.Getenv("MYSQL_USER_PW") + "@tcp(" + os.Getenv("MYSQL_SERVER_URL") + ":" + os.Getenv("MYSQL_SERVER_PORT") + ")/" + os.Getenv("MYSQL_SELECT_DB_NAME") + "?charset=utf8"
	// logger Init
	log.LogInit()
	DB, err = gorm.Open("mysql", DbStr)

	if err != nil {
		log.Logger.Error.Println(err)
	}

	// The table name is automatically named as the plural of the structure name.
	// The column name has the name of the lowercase snake case of the structure field.
	DB.AutoMigrate(&ApiUsage{})

	DB.DB().SetMaxIdleConns(10)  // idle connection pool
	DB.DB().SetMaxOpenConns(100) // Set the maximum number of open connections to the database
}
