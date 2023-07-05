package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

// AbciInfo godoc
//
//	@Summary		Get info about the application
//	@Description	Get info about the application. Upon success, the Cache-Control header will be set with the default maximum age.
//	@Tags			ABCI
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	model.RpcAbciInfo
//	@Failure		500	{object}	model.RpcError
//	@Router			/abci_info [get]
func AbciInfo(c *gin.Context) {
	getPath := strings.Split(c.Request.URL.String(), "/rpc")[1]
	callUrl := os.Getenv("RPC_URL") + getPath
	response, err := http.Get(callUrl)

	if err != nil {
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
