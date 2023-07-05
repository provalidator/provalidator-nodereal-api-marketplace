package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
)

// ConsensusState godoc
//
//	@Summary		Get consensus state
//	@Description	Get consensus state. Not safe to call from inside the ABCI application during a block execution.
//	@Tags			Info
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	model.RpcConsensusState
//	@Failure		500	{object}	model.RpcError
//	@Router			/consensus_state [get]
func ConsensusStatus(c *gin.Context) {
	callUrl := os.Getenv("RPC_URL") + "/consensus_state"
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
