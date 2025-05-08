package keeper

import (
	"context"

	"github.com/qlabs-xyz/qcore/x/pool/types"
)

func (k Keeper) SetTotalMinted(ctx context.Context, minted types.Minted) {
	store := k.storeService.OpenKVStore(ctx)
	b := k.cdc.MustMarshal(&minted)
	store.Set(types.GetMintKey("total_minted_key"), b)
}

func (k Keeper) GetTotalMinted(ctx context.Context) (minted types.Minted, found bool) {
	store := k.storeService.OpenKVStore(ctx)
	mintedKey := types.GetMintKey("total_minted_key")
	b, err := store.Get(mintedKey)

	if b == nil || err != nil {
		return minted, false
	}

	k.cdc.MustUnmarshal(b, &minted)
	return minted, true
}
