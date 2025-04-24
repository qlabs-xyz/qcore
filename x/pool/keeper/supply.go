package keeper

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/qlabs-xyz/qcore/x/pool/types"
)

func (k Keeper) TotalSupply(ctx context.Context) (string, error) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	iterator := storetypes.KVStorePrefixIterator(store, types.TotalSupplyKey)
	defer iterator.Close()

	var supply string
	for ; iterator.Valid(); iterator.Next() {
		supply = string(iterator.Value())
	}
	return supply, nil
}
