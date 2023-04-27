package routes

import (
	"net/http"
	"os"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/provalidator-nodereal-api-marketplace/log"
	lcdRoutes "github.com/provalidator-nodereal-api-marketplace/routes/lcd"
	"github.com/provalidator-nodereal-api-marketplace/util"
)

func Lcd() http.Handler {
	r := setupLcdRouter()
	log.Logger.Info.Println("LCD server is on! port : ", os.Getenv("LCD_PORT"))
	return r
}

//https://docs.tendermint.com/v0.34/rpc/#/Info/validators
//https://v1.cosmos.network/rpc/v0.41.4
func setupLcdRouter() *gin.Engine {
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

	// LCD
	lcd := r.Group("")
	{
		// Gaia REST
		lcd.GET("/node_info", mw, lcdRoutes.Get)

		// Tendermint RPC
		lcd.GET("/syncing", mw, lcdRoutes.Get)
		lcd.GET("/blocks/latest", mw, lcdRoutes.Get)
		lcd.GET("/blocks/:height", mw, lcdRoutes.Get)
		lcd.GET("/validatorsets/latest", mw, lcdRoutes.Get)
		lcd.GET("/validatorsets/:height", mw, lcdRoutes.Get)

		// Transactions
		lcd.POST("/txs", mw, lcdRoutes.Post)
		lcd.POST("/staking/delegators/:delegatorAddr/delegations", mw, lcdRoutes.Post)
		lcd.POST("/staking/delegators/:delegatorAddr/unbonding_delegations", mw, lcdRoutes.Post)
	}
	lcdCosmos := r.Group("/cosmos")
	{
		// Auth
		lcdCosmos.GET("/auth/v1beta1/accounts/:address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/auth/v1beta1/params", mw, lcdRoutes.Get)

		// Bank
		lcdCosmos.GET("/bank/v1beta1/balances/:address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/balances/:address/:denom", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/supply", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/supply/:denom", mw, lcdRoutes.Get)

		// Distribution
		lcdCosmos.GET("/distribution/v1beta1/community_pool", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/delegators/:delegator_address/rewards", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/delegators/:delegator_address/rewards/:validator_address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/delegators/:delegator_address/validators", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/delegators/:delegator_address/withdraw_address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/validators/:validator_address/commission", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/validators/:validator_address/outstanding_rewards", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/validators/:validator_address/slashes", mw, lcdRoutes.Get)

		// Evidence
		lcdCosmos.GET("/evidence/v1beta1/evidence", mw, lcdRoutes.Get)
		lcdCosmos.GET("/evidence/v1beta1/evidence/:evidence_hash", mw, lcdRoutes.Get)

		// Gov
		lcdCosmos.GET("/gov/v1beta1/params/:params_type", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/deposits", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/deposits/:depositor", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/tally", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/votes", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/votes/:voter", mw, lcdRoutes.Get)

		// Mint
		lcdCosmos.GET("/mint/v1beta1/annual_provisions", mw, lcdRoutes.Get)
		lcdCosmos.GET("/mint/v1beta1/inflation", mw, lcdRoutes.Get)
		lcdCosmos.GET("/mint/v1beta1/params", mw, lcdRoutes.Get)

		// Params
		lcdCosmos.GET("/params/v1beta1/params", mw, lcdRoutes.Get)

		// Slashing
		lcdCosmos.GET("/slashing/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/slashing/v1beta1/signing_infos", mw, lcdRoutes.Get)
		lcdCosmos.GET("/slashing/v1beta1/signing_infos/:cons_address", mw, lcdRoutes.Get)

		// Staking
		lcdCosmos.GET("/staking/v1beta1/delegations/:delegator_addr", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegators/:delegator_addr/redelegations", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegators/:delegator_addr/unbonding_delegations", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegators/:delegator_addr/validators", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegators/:delegator_addr/validators/:validator_addr", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/historical_info/:height", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/pool", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr/delegations", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr/delegations/:delegator_addr", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr/delegations/:delegator_addr/unbonding_delegation", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr/unbonding_delegations", mw, lcdRoutes.Get)

		// Upgrade
		lcdCosmos.GET("/upgrade/v1beta1/applied_plan/:name", mw, lcdRoutes.Get)
		lcdCosmos.GET("/upgrade/v1beta1/current_plan", mw, lcdRoutes.Get)
		lcdCosmos.GET("/upgrade/v1beta1/upgraded_consensus_state/:last_height", mw, lcdRoutes.Get)

		// Service
		lcdCosmos.GET("/base/tendermint/v1beta1/blocks/latest", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/blocks/:height", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/node_info", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/syncing", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/validatorsets/latest", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/validatorsets/:height", mw, lcdRoutes.Get)
		lcdCosmos.POST("/tx/v1beta1/simulate", mw, lcdRoutes.Post)
		lcdCosmos.GET("/tx/v1beta1/txs", mw, lcdRoutes.Get)
		lcdCosmos.POST("/tx/v1beta1/txs", mw, lcdRoutes.Post)
		lcdCosmos.GET("/tx/v1beta1/txs/:hash", mw, lcdRoutes.Get)
	}
	lcdIBC := r.Group("/ibc")
	{
		// IBC
		lcdIBC.GET("/core/channel/v1beta1/channels", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/client_state", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/consensus_state/revision/:revision_number/height/:revision_height", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/next_sequence", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/packet_acknowledgements", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/packet_acks/:sequence", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/packet_commitments", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/packet_commitments/:sequence/unreceived_acks", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/packet_commitments/:sequence/unreceived_packets", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/packet_commitments/:sequence", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/channels/:channel_id/ports/:port_id/packet_receipts/:sequence", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1beta1/connections/:connection/channels", mw, lcdRoutes.Get)
		lcdIBC.GET("/client/v1beta1/params", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1beta1/client_states", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1beta1/client_states/:client_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1beta1/consensus_states/:client_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1beta1/consensus_states/:client_id/revision/:revision_number/height/:revision_height", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1beta1/client_connections/:client_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1beta1/connections", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1beta1/connections/:connection_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1beta1/connections/:connection_id/client_state", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1beta1/connections/:connection_id/consensus_state/revision/:revision_number/height/:revision_height", mw, lcdRoutes.Get)
		lcdIBC.GET("/applications/transfer/v1beta1/denom_traces", mw, lcdRoutes.Get)
		lcdIBC.GET("/applications/transfer/v1beta1/denom_traces/:hash", mw, lcdRoutes.Get)
		lcdIBC.GET("/applications/transfer/v1beta1/params", mw, lcdRoutes.Get)

	}
	return r
}
