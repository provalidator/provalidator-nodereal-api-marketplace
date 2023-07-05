package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

// DumpConsensusStatus godoc
//
//	@Summary		Get consensus state
//	@Description	Get consensus state. Not safe to call from inside the ABCI application during a block execution.
//	@Tags			Info
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	model.RpcDumpConsensusState
//	@Failure		500	{object}	model.RpcError
//	@Router			/dump_consensus_state [get]
func DumpConsensusStatus(c *gin.Context) {
	callUrl := os.Getenv("RPC_URL") + "/dump_consensus_state"
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
