package wasmbinding

import (
	"encoding/json"

	errortypes "github.com/qlabs-xyz/qcore/errors"
	"github.com/qlabs-xyz/qcore/wasmbinding/bindings"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.QCoreQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrapf(errortypes.ErrInvalidType, "[CustomQuerier][Unmarshal Contract Query Result] failed. Contract query is not valid, couldn't be parsed.")
		}

		switch {
		// case contractQuery.Minters != nil:

		// 	minter, err := GetMinter(ctx, *qp.keeper)
		// 	if err != nil {
		// 		return nil, sdkerrors.Wrap(errortypes.ErrInvalidType, "[CustomQuerier][GetMinter] failed.")
		// 	}

		// 	response := bindings.MintersResponse{AnnualProvisions: minter.AnnualProvisions, CurrentEpoch: minter.CurrentEpoch, Identifier: minter.Identifier}

		// 	bz, err := json.Marshal(response)
		// 	if err != nil {
		// 		return nil, sdkerrors.Wrapf(errortypes.ErrInvalidType, "[CustomQuerier][Marshal] failed to marshal response")
		// 	}

		// 	return bz, nil

		default:
			return nil, sdkerrors.Wrapf(errortypes.ErrInvalidType, "[CustomQuerier] failed. unknown qcore chain query variante.")
		}
	}
}

// func GetMinter(ctx sdk.Context, pool poolKeepers.Keeper) (bindings.MintersResponse, error) {

// 	log.Println("############## Smart contract query for fetching minter is Started ##############")

// 	var response bindings.MintersResponse

// 	logger := gemmint.Logger(ctx)

// 	minter, err := gemmint.GetAllMinter(ctx)
// 	if err != nil {

// 	}

// 	if logger != nil {
// 		logger.Info("Fetching smart contract query for minter successfully done.",
// 			"query", "GetAllMinter",
// 		)
// 	}

// 	response.AnnualProvisions = strconv.FormatUint(minter[0].AnnualProvisions, 10)
// 	response.CurrentEpoch = strconv.FormatUint(uint64(minter[0].CurrentEpoch), 10)
// 	response.Identifier = minter[0].Identifier

// 	log.Println("############## End of Smart contract query for fetching minter ##############")

// 	return response, nil
// }
