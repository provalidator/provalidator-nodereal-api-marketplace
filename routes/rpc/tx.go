package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

type TxBind struct {
	Hash string `form:"hash" binding:"required"`
	Prov string `form:"prov,default=true"`
}

// Tx godoc
//
//	@Summary		Node heartbeat
//	@Description	Get Transaction
//	@Tags			Tx
//	@Accept			*/*
//	@Produce		json
//	@Param			hash	query		string	true	"Transaction hash (0x + Hash)"	default(0xAAAE86DA167A62EF57F30B857E9F37FE4E13060FE2569BB700FB286D03F1F454)
//	@Param			prov	query		boolean	false	"Include proofs of the transaction's inclusion in the block"	default(false)
//	@Success		200		{object}	model.RpcTx
//	@Failure		500		{object}	model.RpcError
//	@Router			/tx [get]
func Tx(c *gin.Context) {
	// Validation
	req := &TxBind{}
	chk := c.BindQuery(req)

	if chk != nil {
		c.JSON(http.StatusBadRequest, chk.Error())
		return
	}

	// Add 0x
	if !strings.HasPrefix(req.Hash, "0x") {
		req.Hash = "0x" + req.Hash
	}

	// Make URL
	getPath := "/tx"
	queryParams := "?hash=" + req.Hash + "&prov=" + req.Prov
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
