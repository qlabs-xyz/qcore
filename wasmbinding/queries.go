package wasmbinding

import (
	keeper "github.com/qlabs-xyz/qcore/x/pool/keeper"
)

type QueryPlugin struct {
	keeper *keeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(
	pool *keeper.Keeper,
) *QueryPlugin {
	return &QueryPlugin{
		keeper: pool,
	}
}
