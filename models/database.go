package models

import (
	"os"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/provalidator-nodereal-api-marketplace/log"
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

func WriteLog(sub string, token string, ip string, method string) {
	//log.Logger.Info.Println("WriteLog. Token:", token, ", IP:", ip, ", Method:", method)
	now := time.Now()
	date := now.Format("2006-01-02 15:04:05")
	input := ApiUsage{Sub: sub, Token: token, Ip: ip, Method: method, Date: date, ChainName: os.Getenv("CHAIN_NAME")}
	DB.Create(&input)

	// Usage count
	ym := now.Format("0601")
	uqKey := ym + "-" + sub
	DB.Exec("INSERT INTO nodereal.api_usages_counts(uqkey,chain) values (?,?) ON DUPLICATE KEY UPDATE uqkey = ?, counts = counts+1;", uqKey, "cosmos", uqKey)

}

func Count(sub string) int {
	apiUsagesCount := ApiUsagesCount{}
	now := time.Now()
	// Usage count
	ym := now.Format("0601")
	uqKey := ym + "-" + sub
	DB.Raw("SELECT counts FROM nodereal.api_usages_counts WHERE uqkey = ?", uqKey).Scan(&apiUsagesCount)
	return apiUsagesCount.Counts
}
