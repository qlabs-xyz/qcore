package keeper

import (
	"github.com/qlabs-xyz/qcore/x/pool/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {

	// Set if defined
	// for _, minter := range genState.Minters {
	// 	k.SetMinter(ctx, minter)
	// }

	// k.SetParams(ctx, genState.Params)
}
