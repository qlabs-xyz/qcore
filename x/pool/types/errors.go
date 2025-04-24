package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrUnauthorizedMinter = sdkerrors.Register(ModuleName, 1, "unauthorized minter")
	ErrInsufficientPool   = sdkerrors.Register(ModuleName, 2, "insufficient pool balance")
	ErrInvalidBlock       = sdkerrors.Register(ModuleName, 3, "invalid block height")
)
