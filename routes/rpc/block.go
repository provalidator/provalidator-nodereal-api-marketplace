package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

type BlockBind struct {
	Height string `form:"height"`
}

// Block godoc
//
//	@Summary		Get block at a specified height
//	@Description	Get Block. If the height field is set to a non-default value, upon success, the Cache-Control header will be set with the default maximum age.
//	@Tags			Info
//	@Accept			*/*
//	@Produce		json
//	@Param			height	query		int	false	"height to return. If no height is provided, it will fetch validator set which corresponds to the latest block."	default(14992000)
//	@Success		200		{object}	model.RpcBlock
//	@Failure		500		{object}	model.RpcError
//	@Router			/block [get]
func Block(c *gin.Context) {
	// Validation
	req := &BlockBind{}
	chk := c.BindQuery(req)

	if chk != nil {
		c.JSON(http.StatusBadRequest, chk.Error())
		return
	}

	getPath := "/block"
	queryParams := "?height=" + req.Height
	callUrl := os.Getenv("RPC_URL") + getPath + queryParams
	response, err := http.Get(callUrl)

	if err != nil {
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
