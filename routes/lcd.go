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

	// Ethermint
	lcdEthermint := r.Group("/ethermint")
	{
		lcdEthermint.GET("/evm/v1/account/:address", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/balances/:address", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/base_fee", mw, lcdRoutes.Get)
		lcdEthermint.GET("/feemarket/v1/block_gas", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/codes/:address", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/cosmos_account/:address", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/estimate_gas", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/eth_call", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/params", mw, lcdRoutes.Get)
		lcdEthermint.GET("/feemarket/v1/base_fee", mw, lcdRoutes.Get)
		lcdEthermint.GET("/feemarket/v1/params", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/storage/:address/:key", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/trace_block", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/trace_tx", mw, lcdRoutes.Get)
		lcdEthermint.GET("/evm/v1/validator_account/:cons_address", mw, lcdRoutes.Get)
	}
	// Cosmos
	lcdCosmos := r.Group("/cosmos")
	{
		lcdCosmos.GET("/auth/v1beta1/accounts", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/balances/:address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/evidence/v1beta1/evidence", mw, lcdRoutes.Get)
		lcdCosmos.GET("/feegrant/v1beta1/allowance/:granter/:grantee", mw, lcdRoutes.Get)
		lcdCosmos.GET("/feegrant/v1beta1/allowances/:grantee", mw, lcdRoutes.Get)
		lcdCosmos.GET("/auth/v1beta1/accounts/:address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/auth/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/balances/:address/by_denom", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/community_pool", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr/delegations/:delegator_addr", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/delegators/:delegator_address/rewards/:validator_address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/delegators/:delegator_address/rewards", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegations/:delegator_addr", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegators/:delegator_addr/unbonding_delegations", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegators/:delegator_addr/validators/:validator_addr", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegators/:delegator_addr/validators", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/delegators/:delegator_address/withdraw_address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/denoms_metadata/:denom", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/denoms_metadata", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/deposits/:depositor", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/deposits", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/delegators/:delegator_address/validators", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/evidence/v1beta1/evidence/:evidence_hash", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/params/:params_type", mw, lcdRoutes.Get)
		lcdCosmos.GET("/authz/v1beta1/grants/grantee/:grantee", mw, lcdRoutes.Get)
		lcdCosmos.GET("/authz/v1beta1/grants/granter/:granter", mw, lcdRoutes.Get)
		lcdCosmos.GET("/authz/v1beta1/grants", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/historical_info/:height", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/pool", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/delegators/:delegator_addr/redelegations", mw, lcdRoutes.Get)
		lcdCosmos.GET("/slashing/v1beta1/signing_infos/:cons_address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/slashing/v1beta1/signing_infos", mw, lcdRoutes.Get)
		lcdCosmos.GET("/slashing/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/spendable_balances/:address", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/params", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/supply/:denom", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/tally", mw, lcdRoutes.Get)
		lcdCosmos.GET("/bank/v1beta1/supply", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr/delegations/:delegator_addr/unbonding_delegation", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/validators/:validator_address/commission", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr/delegations", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/validators/:validator_address/outstanding_rewards", mw, lcdRoutes.Get)
		lcdCosmos.GET("/distribution/v1beta1/validators/:validator_address/slashes", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators/:validator_addr/unbonding_delegations", mw, lcdRoutes.Get)
		lcdCosmos.GET("/staking/v1beta1/validators", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/votes/:voter", mw, lcdRoutes.Get)
		lcdCosmos.GET("/gov/v1beta1/proposals/:proposal_id/votes", mw, lcdRoutes.Get)

		// Cosmos - Service
		lcdCosmos.POST("/tx/v1beta1/txs", mw, lcdRoutes.Post)
		lcdCosmos.GET("/base/tendermint/v1beta1/blocks/:height", mw, lcdRoutes.Get)
		lcdCosmos.GET("/tx/v1beta1/txs/block/:height", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/blocks/latest", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/validatorsets/latest", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/node_info", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/syncing", mw, lcdRoutes.Get)
		lcdCosmos.GET("/tx/v1beta1/txs/:hash", mw, lcdRoutes.Get)
		lcdCosmos.GET("/tx/v1beta1/txs", mw, lcdRoutes.Get)
		lcdCosmos.GET("/base/tendermint/v1beta1/validatorsets/:height", mw, lcdRoutes.Get)
		lcdCosmos.POST("/tx/v1beta1/simulate", mw, lcdRoutes.Post)
	}
	// Evmos
	lcdEvmos := r.Group("/evmos")
	{
		lcdEvmos.GET("/incentives/v1/allocation_meters/:denom", mw, lcdRoutes.Get)
		lcdEvmos.GET("/incentives/v1/allocation_meters", mw, lcdRoutes.Get)
		lcdEvmos.GET("/vesting/v1/balances/:address", mw, lcdRoutes.Get)
		lcdEvmos.GET("/inflation/v1/circulating_supply", mw, lcdRoutes.Get)
		lcdEvmos.GET("/claims/v1/params", mw, lcdRoutes.Get)
		lcdEvmos.GET("/claims/v1/claims_records/:address", mw, lcdRoutes.Get)
		lcdEvmos.GET("/claims/v1/claims_records", mw, lcdRoutes.Get)
		lcdEvmos.GET("/epochs/v1/current_epoch", mw, lcdRoutes.Get)
		lcdEvmos.GET("/erc20/v1/params", mw, lcdRoutes.Get)
		lcdEvmos.GET("/epochs/v1/epochs", mw, lcdRoutes.Get)
		lcdEvmos.GET("/inflation/v1/epoch_mint_provision", mw, lcdRoutes.Get)
		lcdEvmos.GET("/incentives/v1/gas_meters/:contract/:participant", mw, lcdRoutes.Get)
		lcdEvmos.GET("/incentives/v1/gas_meters/:contract", mw, lcdRoutes.Get)
		lcdEvmos.GET("/incentives/v1/incentives/:contract", mw, lcdRoutes.Get)
		lcdEvmos.GET("/incentives/v1/incentives", mw, lcdRoutes.Get)
		lcdEvmos.GET("/incentives/v1/params", mw, lcdRoutes.Get)
		lcdEvmos.GET("/inflation/v1/params", mw, lcdRoutes.Get)
		lcdEvmos.GET("/inflation/v1/inflation_rate", mw, lcdRoutes.Get)
		lcdEvmos.GET("/inflation/v1/period", mw, lcdRoutes.Get)
		lcdEvmos.GET("/recovery/v1/params", mw, lcdRoutes.Get)
		lcdEvmos.GET("/inflation/v1/skipped_epochs", mw, lcdRoutes.Get)
		lcdEvmos.GET("/erc20/v1/token_pairs/:token", mw, lcdRoutes.Get)
		lcdEvmos.GET("/erc20/v1/token_pairs", mw, lcdRoutes.Get)
		lcdEvmos.GET("/claims/v1/total_unclaimed", mw, lcdRoutes.Get)
	}

	// IBC
	lcdIBC := r.Group("/ibc")
	{
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/client_state", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/consensus_state/revision/:revision_number/height/:revision_height", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1/client_connections/:client_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/client/v1/params", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1/client_states/:client_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1/client_states", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1/client_status/:client_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1/connections/:connection_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/connections/:connection/channels", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1/connections/:connection_id/client_state", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1/connections/:connection_id/consensus_state/revision/:revision_number/height/:revision_height", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/connection/v1/connections", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1/consensus_states/:client_id/revision/:revision_number/height/:revision_height", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1/consensus_states/:client_id", mw, lcdRoutes.Get)
		lcdIBC.GET("/apps/transfer/v1/denom_hashes/:trace", mw, lcdRoutes.Get)
		lcdIBC.GET("/apps/transfer/v1/denom_traces/:hash", mw, lcdRoutes.Get)
		lcdIBC.GET("/apps/transfer/v1/denom_traces", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/next_sequence", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/packet_acks/:sequence", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/packet_acknowledgements", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/packet_commitments/:sequence", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/packet_commitments", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/packet_receipts/:sequence", mw, lcdRoutes.Get)
		lcdIBC.GET("/apps/transfer/v1/params", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/packet_commitments/:sequence/unreceived_acks", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/channel/v1/channels/:channel_id/ports/:port_id/packet_commitments/:sequence/unreceived_packets", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1/upgraded_client_states", mw, lcdRoutes.Get)
		lcdIBC.GET("/core/client/v1/upgraded_consensus_states", mw, lcdRoutes.Get)

	}
	return r
}
