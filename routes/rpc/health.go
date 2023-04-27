package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Health godoc
//	@Summary		Node heartbeat
//	@Description	Get node health. Returns empty result (200 OK) on success, no response - in case of an error.
//	@Tags			Info
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	model.RpcHealth
//	@Failure		500	{object}	model.RpcError
//	@Router			/health [get]
func Health(c *gin.Context) {
	getPath := "/health"
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
