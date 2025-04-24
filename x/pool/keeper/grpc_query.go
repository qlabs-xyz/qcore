package keeper

import (
	"github.com/qlabs-xyz/qcore/x/pool/types"
)

type QueryServer struct {
	Keeper
}

var _ types.QueryServer = Keeper{}

func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &QueryServer{
		Keeper: keeper,
	}
}
