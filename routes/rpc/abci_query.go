package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

type AbciQueryBind struct {
	Path   string `form:"path" binding:"required"`
	Data   string `form:"data" binding:"required"`
	Height string `form:"height,default=0"`
	Prov   string `form:"prov,default=true"`
}

// AbciQuery godoc
//
//	@Summary		Query the application for some information.
//	@Description	Query the application for some information.
//	@Tags			ABCI
//	@Accept			*/*
//	@Produce		json
//	@Param			path		query		string	true	"Path to the data ("/a/b/c")"	default(/a/b/c)
//	@Param			data		query		string	true	"Data"	default(IHAVENOIDEA)
//	@Param			height		query		integer	false	"Height (0 means latest)"	default(1)
//	@Param			prov		query		boolean	false	"Include proofs of the transaction's inclusion in the block"	default(true)
//	@Success		200	{object}	model.AbciQuery
//	@Failure		500	{object}	model.RpcError
//	@Router			/abci_query [get]
func AbciQuery(c *gin.Context) {
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
