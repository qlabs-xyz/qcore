package keeper

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qlabs-xyz/qcore/x/pool/constants"
	"github.com/qlabs-xyz/qcore/x/pool/types"
)

func (k Keeper) GetBlockEmission(goCtx context.Context, req *types.QueryBlockEmissionRequest) (*types.QueryBlockEmissionResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	if req == nil {
		return nil, errors.New("request is nil")
	}

	fmt.Println("ctx.BlockHeight( --------------->", ctx.BlockHeight())

	if ctx.BlockHeight() < constants.TransitionBlockNumber {

		if req.BlockNumber < 0 {
			return nil, errors.New("blocknumber is 0")
		}

		tokens, err := k.CalculateExponentialTokens(req.BlockNumber)
		if err != nil {
			return nil, errors.New("CalculateExponentialTokens failed")
		}
		return &types.QueryBlockEmissionResponse{Tokens: tokens}, nil
	}

	tokens, err := k.CalculateFixedTokens(goCtx)
	if err != nil {
		return nil, errors.New("CalculateFixedTokens failed")
	}

	return &types.QueryBlockEmissionResponse{Tokens: tokens}, nil
}
