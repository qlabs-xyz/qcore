package keeper

import (
	"github.com/qlabs-xyz/qcore/x/pool/types"
)

type msgServer struct {
	Keeper
}

/** NewMsgServerImpl returns an implementation of the MsgServer interface for the provided Keeper. */
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
