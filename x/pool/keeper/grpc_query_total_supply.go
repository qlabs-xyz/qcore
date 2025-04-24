package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qlabs-xyz/qcore/x/pool/types"
)

func (k Keeper) GetTotalSupply(goCtx context.Context, req *types.QueryTotalSupplyRequest) (*types.QueryTotalSupplyResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	totalSupply, err := k.TotalSupply(ctx)
	if err != nil {
		return nil, errors.New("fetching total supply failed")
	}

	return &types.QueryTotalSupplyResponse{TotalSupply: totalSupply}, nil
}
