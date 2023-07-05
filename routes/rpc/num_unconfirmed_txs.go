package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

//	 NumUnconfirmedTxs godoc
//		@Summary		Get data about unconfirmed transactions
//		@Description	Get data about unconfirmed transactions
//		@Tags			Info
//		@Accept			*/*
//		@Produce		json
//		@Success		200	{object}	model.NumUnconfirmedTxs
//		@Failure		500	{object}	model.RpcError
//		@Router			/num_unconfirmed_txs [get]
func NumUnconfirmedTxs(c *gin.Context) {
	getPath := "/num_unconfirmed_txs"
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
