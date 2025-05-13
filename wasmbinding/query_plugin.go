package wasmbinding

import (
	"encoding/json"

	errortypes "github.com/outbe/outbe-node/errors"
	"github.com/outbe/outbe-node/wasmbinding/bindings"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.QcoreQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrapf(errortypes.ErrInvalidType, "[CustomQuerier][Unmarshal Contract Query Result] failed. Contract query is not valid, couldn't be parsed.")
		}

		switch {

		default:
			return nil, sdkerrors.Wrapf(errortypes.ErrInvalidType, "[CustomQuerier] failed. unknown qcore chain query variante.")
		}
	}
}
