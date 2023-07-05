package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

func Get(c *gin.Context) {
	getPath := c.Request.URL.String()
	callUrl := os.Getenv("LCD_URL") + getPath

	log.Logger.Info.Println(callUrl)

	response, err := http.Get(callUrl)

	if err != nil {
		log.Logger.Error.Println(err.Error())
		c.Status(http.StatusServiceUnavailable)
		return
	}
	err2 := response.Body.Close()

	if err2 != nil {
		log.Logger.Error.Println(err.Error())
	}

	var v interface{}
	json.NewDecoder(response.Body).Decode(&v)
	c.JSON(200, v) // Write Body
}
