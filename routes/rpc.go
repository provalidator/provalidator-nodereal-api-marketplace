package routes

import (
	"net/http"
	"os"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
	rpcRoutes "github.com/provalidator-nodereal-api-marketplace/routes/rpc"
	"github.com/provalidator-nodereal-api-marketplace/util"
)

func Rpc() http.Handler {
	r := setupRpcRouter()
	log.Logger.Info.Println("RPC server is on! port : ", os.Getenv("RPC_PORT"))
	return r
}

//https://docs.tendermint.com/v0.34/rpc/#/Info/validators
//https://v1.cosmos.network/rpc/v0.41.4
func setupRpcRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	// Rate Limit
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: util.ErrorHandler,
		KeyFunc:      util.KeyFunc,
	})

	// Middleware
	r.Use(AuthorizationMiddleware)

	// RPC
	rpc := r.Group("")
	{
		// Info
		rpc.GET("/health", mw, rpcRoutes.Health)
		rpc.GET("/net_info", mw, rpcRoutes.NetInfo)
		rpc.GET("/status", mw, rpcRoutes.Status)
		rpc.GET("/consensus_state", mw, rpcRoutes.ConsensusStatus)
		rpc.GET("/dump_consensus_state", mw, rpcRoutes.DumpConsensusStatus)
		rpc.GET("/consensus_params", mw, rpcRoutes.ConsensusParams)
		rpc.GET("/tx", mw, rpcRoutes.Tx)
		rpc.GET("/check_tx", mw, rpcRoutes.CheckTx)
		rpc.GET("/block", mw, rpcRoutes.Block)
		rpc.GET("/validators", mw, rpcRoutes.Validators)
		rpc.GET("/unconfirmed_txs", mw, rpcRoutes.UnconfirmedTxs)
		rpc.GET("/num_unconfirmed_txs", mw, rpcRoutes.NumUnconfirmedTxs)

		// ABCI
		rpc.GET("/abci_info", mw, rpcRoutes.AbciInfo)
		rpc.GET("/abci_query", mw, rpcRoutes.AbciQuery)
	}

	return r
}
