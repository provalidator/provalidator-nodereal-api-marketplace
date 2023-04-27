package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//  NumUnconfirmedTxs godoc
//	@Summary		Get data about unconfirmed transactions
//	@Description	Get data about unconfirmed transactions
//	@Tags			Info
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	model.NumUnconfirmedTxs
//	@Failure		500	{object}	model.RpcError
//	@Router			/num_unconfirmed_txs [get]
func NumUnconfirmedTxs(c *gin.Context) {
	getPath := "/num_unconfirmed_txs"
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
