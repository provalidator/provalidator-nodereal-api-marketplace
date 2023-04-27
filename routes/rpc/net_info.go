package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// NetInfo godoc
//	@Summary		Network informations
//	@Description	Get network info
//	@Tags			Info
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	model.RpcNetInfo
//	@Failure		500	{object}	model.RpcError
//	@Router			/net_info [get]
func NetInfo(c *gin.Context) {
	getPath := "/net_info"
	callUrl := os.Getenv("RPC_URL") + getPath
	response, err := http.Get(callUrl)

	if err != nil {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	var v interface{}
	json.NewDecoder(response.Body).Decode(&v)
	c.JSON(200, v) // Write Body
}
