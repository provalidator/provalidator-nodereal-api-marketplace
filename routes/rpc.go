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

// https://docs.tendermint.com/v0.34/rpc/#/Info/validators
// https://v1.cosmos.network/rpc/v0.41.4
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
		//JSONRPC/HTTP
		rpc.POST("", mw, rpcRoutes.Post)
		rpc.POST("clientVersion", mw, rpcRoutes.Post)
		rpc.POST("sha3", mw, rpcRoutes.Post)
		rpc.POST("version", mw, rpcRoutes.Post)
		rpc.POST("listening", mw, rpcRoutes.Post)
		rpc.POST("peerCount", mw, rpcRoutes.Post)
		rpc.POST("protocolVersion", mw, rpcRoutes.Post)
		rpc.POST("syncing", mw, rpcRoutes.Post)
		rpc.POST("coinbase", mw, rpcRoutes.Post)
		rpc.POST("gasPrice", mw, rpcRoutes.Post)
		rpc.POST("accounts", mw, rpcRoutes.Post)
		rpc.POST("blockNumber", mw, rpcRoutes.Post)
		rpc.POST("getBalance", mw, rpcRoutes.Post)
		rpc.POST("getStorageAt", mw, rpcRoutes.Post)
		rpc.POST("getBlockTransactionCount", mw, rpcRoutes.Post)
		rpc.POST("getBlockTransactionCountByHash", mw, rpcRoutes.Post)
		rpc.POST("getBlockTransactionCountByNumber", mw, rpcRoutes.Post)
		rpc.POST("getCode", mw, rpcRoutes.Post)
		rpc.POST("sign", mw, rpcRoutes.Post)
		rpc.POST("sendTransaction", mw, rpcRoutes.Post)
		rpc.POST("senRawTransaction", mw, rpcRoutes.Post)
		rpc.POST("call", mw, rpcRoutes.Post)
		rpc.POST("estimateGas", mw, rpcRoutes.Post)
		rpc.POST("getBlockByHash", mw, rpcRoutes.Post)
		rpc.POST("getBlockSignersByNumber", mw, rpcRoutes.Post)
		rpc.POST("getBlockSignersByHash", mw, rpcRoutes.Post)
		rpc.POST("getBlockFinalityByNumber", mw, rpcRoutes.Post)
		rpc.POST("getBlockFinalityByHash", mw, rpcRoutes.Post)
		rpc.POST("getCandidates", mw, rpcRoutes.Post)
		rpc.POST("getCandidateStatus", mw, rpcRoutes.Post)
		rpc.POST("getTransactionByHash", mw, rpcRoutes.Post)
		rpc.POST("getTransactionByBlockHashAndIndex", mw, rpcRoutes.Post)
		rpc.POST("getTransactionByBlockNumberAndIndex", mw, rpcRoutes.Post)
		rpc.POST("getTransactionReceipt", mw, rpcRoutes.Post)
		rpc.POST("getAskTree", mw, rpcRoutes.Post)
		rpc.POST("getAsks", mw, rpcRoutes.Post)
		rpc.POST("getBestAsk", mw, rpcRoutes.Post)
		rpc.POST("getBidTree", mw, rpcRoutes.Post)
		rpc.POST("getBids", mw, rpcRoutes.Post)
		rpc.POST("getBestBid", mw, rpcRoutes.Post)
		rpc.POST("getBorrows", mw, rpcRoutes.Post)
		rpc.POST("getInvest", mw, rpcRoutes.Post)
		rpc.POST("getNetworkInformation", mw, rpcRoutes.Post)
	}

	return r
}
