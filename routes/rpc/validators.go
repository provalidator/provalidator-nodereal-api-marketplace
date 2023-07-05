package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

type ValidatorsBind struct {
	Height  string `form:"height"`
	Page    string `form:"page"`
	PerPage string `form:"per_page"`
}

// Validators godoc
//
//	@Summary		Get validator set at a specified height
//	@Description	Get Validators. Validators are sorted by voting power. If the height field is set to a non-default value, upon success, the Cache-Control header will be set with the default maximum age.
//	@Tags			Info
//	@Accept			*/*
//	@Produce		json
//	@Param			height		query		int	false	"height to return. If no height is provided, it will fetch validator set which corresponds to the latest block."	default(14992000)
//	@Param			page		query		int	false	"Page number (1-based)"																								default(1)
//	@Param			per_page	query		int	false	"Number of entries per page (max: 100)"																				default(100)
//	@Success		200			{object}	model.RpcValidators
//	@Failure		500			{object}	model.RpcError
//	@Router			/validators [get]
func Validators(c *gin.Context) {
	// Validation
	req := &ValidatorsBind{}
	chk := c.BindQuery(req)

	if chk != nil {
		c.JSON(http.StatusBadRequest, chk.Error())
		return
	}
	getPath := "/validators"
	queryParams := "?height=" + req.Height + "&page=" + req.Page + "&per_page=" + req.PerPage
	callUrl := os.Getenv("RPC_URL") + getPath + queryParams
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
