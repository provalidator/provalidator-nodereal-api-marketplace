package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

type CheckTxBind struct {
	Tx string `form:"tx" binding:"required"`
}

// CheckTx godoc
//
//	@Summary		Checks the transaction without executing it.
//	@Description	The transaction won't be added to the mempool. Please refer to https://docs.tendermint.com/v0.34/tendermint-core/using-tendermint.html#formatting for formatting/encoding rules. Upon success, the Cache-Control header will be set with the default maximum age.
//	@Tags			Tx
//	@Accept			*/*
//	@Produce		json
//	@Param			tx	query		string	true	"The transaction"	default(0xAAAE86DA167A62EF57F30B857E9F37FE4E13060FE2569BB700FB286D03F1F454)
//	@Success		200		{object}	model.RpcCheckTx
//	@Failure		500		{object}	model.RpcError
//	@Router			/check_tx [get]
func CheckTx(c *gin.Context) {
	// Validation
	req := &CheckTxBind{}
	chk := c.BindQuery(req)

	if chk != nil {
		c.JSON(http.StatusBadRequest, chk.Error())
		return
	}
	// Add 0x
	if !strings.HasPrefix(req.Tx, "0x") {
		req.Tx = "0x" + req.Tx
	}

	// Make URL
	getPath := "/tx"
	queryParams := "?tx=" + req.Tx
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
