package models

import (
	"os"
	"regexp"
	"runtime"
	"testing"

	"github.com/provalidator-nodereal-api-marketplace/config"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

func TestDB(t *testing.T) {

	// Config *
	// logger Init
	log.LogInit()

	// Path
	projectName := regexp.MustCompile(`^(.*` + "provalidator-nodereal-api-marketplace" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := string(projectName.Find([]byte(currentWorkDirectory)))
	os.Setenv("ROOT_PATH", rootPath) //set Root path

	// env
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		config.LoadEnv(rootPath + "/config/local.env")
		log.Logger.Trace.Println("local.env")
	} else {
		config.LoadEnv(rootPath + "/config/prod.env")
		log.Logger.Trace.Println("prod.env")
	}
	log.Logger.Trace.Println("runtime.GOOS", runtime.GOOS)

	log.Logger.Info.Println("TestDB")

	// DB
	ConnectDatabase()
	log.Logger.Info.Println(DbStr)
	Read()
	Write()
	Read()
	Update()
	Read()
	Delete()
	Read()

}

func Read() {
	var apiUsage []ApiUsage
	DB.Find(&apiUsage)
	log.Logger.Info.Println("Read test")
	log.Logger.Info.Println(apiUsage)
}

func Check(token string) {
	var apiUsage = ApiUsage{}
	// DB.Where("token = ?", token).First(apiUsage)
	DB.Exec("").First(apiUsage)
}

func Write() {
	log.Logger.Info.Println("Write test")
	input := ApiUsage{Token: "test", Method: "/asddf", Date: "2012-02-20 12:11:11"}
	DB.Create(&input)
}

func Delete() {
	log.Logger.Info.Println("Delete test")
	delete := ApiUsage{}
	if err := DB.Where("token = ?", "test2").First(&delete).Error; err != nil {
		log.Logger.Error.Println("Delete test err", err)
		return
	}
	DB.Delete(delete)
}

func Update() {
	log.Logger.Info.Println("Update test")
	update := ApiUsage{}
	DB.Model(&update).Update(ApiUsage{Token: "test2"})
}
