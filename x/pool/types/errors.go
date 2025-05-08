package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrUnauthorizedMinter = sdkerrors.Register(ModuleName, 16, "unauthorized minter")
	ErrInsufficientPool   = sdkerrors.Register(ModuleName, 17, "insufficient pool balance")
	ErrInvalidBlock       = sdkerrors.Register(ModuleName, 18, "invalid block height")
)
