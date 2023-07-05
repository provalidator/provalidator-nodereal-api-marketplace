package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

type UnconfirmedTxsBind struct {
	limit string `form:"limit,default=30"`
}

//	 UnconfirmedTxs godoc
//		@Summary		Get list of unconfirmed transactions
//		@Description	Get list of unconfirmed transactions
//		@Tags			Info
//		@Accept			*/*
//		@Produce		json
//		@Param			limit	query		int	false	"Maximum number of unconfirmed transactions to return (max 100)"	default(30)
//		@Success		200		{object}	model.UnconfirmedTxs
//		@Failure		500		{object}	model.RpcError
//		@Router			/unconfirmed_txs [get]
func UnconfirmedTxs(c *gin.Context) {
	// Validation
	req := &UnconfirmedTxsBind{}
	chk := c.BindQuery(req)

	if chk != nil {
		c.JSON(http.StatusBadRequest, chk.Error())
		return
	}

	getPath := "/unconfirmed_txs"
	queryParams := "?limit=" + req.limit
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
