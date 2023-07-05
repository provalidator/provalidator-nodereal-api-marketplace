package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

func Post(c *gin.Context) {
	getPath := c.Request.URL.String()
	callUrl := os.Getenv("RPC_URL") + getPath
	response, err := http.Post(callUrl, "application/json", c.Request.Body)
	if err != nil {
		log.Logger.Error.Println(err.Error())
		c.Status(http.StatusServiceUnavailable)
		return
	}
	var v interface{}
	json.NewDecoder(response.Body).Decode(&v)
	c.JSON(200, v) // Write Body

	err2 := response.Body.Close()

	if err2 != nil {
		log.Logger.Error.Println(err.Error())
	}
}
