package keeper

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/qlabs-xyz/qcore/x/pool/types"
)

func (k Keeper) SetTribute(ctx context.Context, tributeDetails types.Tribute) error {
	store := k.storeService.OpenKVStore(ctx)
	b := k.cdc.MustMarshal(&tributeDetails)
	return store.Set(types.GetTributeKey(tributeDetails.Creator), b)
}

func (k Keeper) GetTribute(ctx context.Context) (list []types.Tribute) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	iterator := storetypes.KVStorePrefixIterator(store, types.TributeKey)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tribute
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}
