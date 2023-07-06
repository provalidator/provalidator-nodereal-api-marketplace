package main

import (
	"crypto/rsa"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"time"

	"github.com/provalidator-nodereal-api-marketplace/config"
	"github.com/provalidator-nodereal-api-marketplace/log"
	"github.com/provalidator-nodereal-api-marketplace/models"
	"github.com/provalidator-nodereal-api-marketplace/routes"
	"golang.org/x/sync/errgroup"
)

var (
	parsedPubKey *rsa.PublicKey
	g            errgroup.Group
)

func main() {
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

	// DB
	models.ConnectDatabase()

	lcdServer := &http.Server{
		Addr:         ":" + os.Getenv("LCD_PORT"),
		Handler:      routes.Lcd(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	rpcServer := &http.Server{
		Addr:         ":" + os.Getenv("RPC_PORT"),
		Handler:      routes.Rpc(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return lcdServer.ListenAndServe()
	})

	g.Go(func() error {
		return rpcServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Logger.Error.Println(err)
	}
}
