package rest

import (
	"github.com/maticnetwork/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
	tmLog "github.com/maticnetwork/tendermint/libs/log"

	"github.com/maticnetwork/heimdall/helper"
)

// RestLogger for staking module logger
var RestLogger tmLog.Logger

func init() {
	RestLogger = helper.Logger.With("module", "staking/rest")
}

// RegisterRoutes registers checkpoint-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
