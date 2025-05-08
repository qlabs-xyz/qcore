package keeper

import (
	"context"
	"errors"
	"strconv"

	errortypes "github.com/qlabs-xyz/qcore/errors"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qlabs-xyz/qcore/app/params"
	"github.com/qlabs-xyz/qcore/x/pool/constants"
	"github.com/qlabs-xyz/qcore/x/pool/types"
)

func (k msgServer) MintTribute(goCtx context.Context, msg *types.MsgMintTributeRequest) (*types.MsgMintTributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	logger.Info("########## Mint Tribute Transaction Started ##########")

	// Retrieve minter details
	// minter, _ := k.GetAllMinter(ctx)
	// if len(minter) == 0 {
	// 	logger.Error("[MintTribute][GetAllMinter] No minter found in state")
	// 	return nil, errors.New("no minter configured in state")
	// }

	// Check if the tribute amount exceeds the available annual provisions
	// if constants.AnnualProvisions < msg.RedemptionAmount {
	// 	logger.Error("[MintTribute] Redemption amount exceeds the annual minting provisions",
	// 		"mint_amount", msg.MintAmount,
	// 		"annual_provisions", constants.AnnualProvisions,
	// 	)
	// 	return nil, errors.New("redemption amount exceeds the minting amount available for the current period")
	// }

	// Validate whitelist and ensure the contract is eligible
	// whitelists := k.GetWhitelist(ctx)
	// if len(whitelists) == 0 {
	// 	logger.Error("[MintTribute][GetWhitelist] No eligible contracts found in the whitelist")
	// 	return nil, errors.New("no eligible contract found in the whitelist")
	// }

	// contractEligible := false
	// for _, contract := range whitelists[0].EligibleContracts {
	// 	if contract.ContractAddress == msg.ContractAddress {
	// 		contractEligible = true
	// 		break
	// 	}
	// }

	// if !contractEligible {
	// 	logger.Error("[MintTribute] Contract not eligible for minting",
	// 		"creator", msg.Creator,
	// 		"contract_address", msg.ContractAddress)
	// 	return nil, errors.New("contract is not eligible for minting")
	// }

	provisioning, _ := k.BlockProvisionAmount(ctx)
	totalMinted, _ := k.GetTotalMinted(ctx)

	totalMintedAmount, _ := strconv.ParseUint(totalMinted.TotalMinted, 10, 64)

	remainProvisioning := provisioning - totalMintedAmount

	if remainProvisioning < msg.MintAmount {
		logger.Error("[MintTribute] Failed to mint coins. No coin for mint.",
			"mint_amount", msg.MintAmount,
			"coin_for_mint", remainProvisioning,
		)
		return nil, sdkerrors.Wrap(errortypes.ErrInvalidRequest, "[MintTribute] failed. No coin for mint.")
	}

	// Mint the coins
	mintCoin := sdk.NewCoin(params.BondDenom, math.NewInt(int64(msg.MintAmount)))
	mintCoins := sdk.NewCoins(mintCoin)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		logger.Error("[MintTribute][MintCoins] Failed to mint coins",
			"mint_amount", msg.MintAmount,
			"error", err,
		)
		return nil, errors.New("failed to mint coins")
	}

	totalMintedAmount += msg.MintAmount
	str := strconv.FormatUint(totalMintedAmount, 10)
	newMinted := types.Minted{
		TotalMinted: str,
	}
	k.SetTotalMinted(ctx, newMinted)

	// Update minter total minted
	// minter[0].TotalMinted += msg.MintAmount
	// k.bankKeeper.SetMinter(ctx, minter[0])

	// Validate the recipient address
	recipientAddress, err := sdk.AccAddressFromBech32(msg.ReceiptAddress)
	if err != nil {
		logger.Error("[MintTribute][AccAddressFromBech32] Invalid recipient address",
			"receipt_address", msg.ReceiptAddress,
			"error", err,
		)
		return nil, errors.New("invalid recipient address")
	}

	// Send minted coins to the recipient
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		recipientAddress,
		sdk.Coins{mintCoin},
	)
	if err != nil {
		logger.Error("[MintTribute][SendCoinsFromModuleToAccount] Failed to send coins",
			"receipt_address", msg.ReceiptAddress,
			"mint_amount", msg.MintAmount,
			"error", err,
		)
		return nil, errors.New("failed to send coins to recipient")
	}

	// Save redemption details
	newRedemption := types.Tribute{
		Creator:          msg.Creator,
		ContractAddress:  msg.ContractAddress,
		RecipientAddress: msg.ReceiptAddress,
		Amount:           msg.MintAmount,
	}
	k.SetTribute(ctx, newRedemption)

	logger.Info("########## Mint Tribute Transaction Transaction Completed ##########",
		"recipient", msg.ReceiptAddress,
		"mint_amount", msg.MintAmount,
	)

	return &types.MsgMintTributeResponse{}, nil
}

func (k msgServer) BlockProvisionAmount(ctx sdk.Context) (uint64, error) {

	if ctx.BlockHeight() < constants.TransitionBlockNumber {

		tokens, err := k.CalculateExponentialBlockEmission(ctx.BlockHeight())
		if err != nil {
			return 0, sdkerrors.Wrapf(errortypes.ErrInvalidRequest, "[BlockProvisionAmount][CalculateExponentialTokens] failed. ")
		}
		val, _ := strconv.ParseUint(tokens, 10, 64)
		return val, nil
	}

	tokens, err := k.CalculateFixedBlockEmission(ctx)
	if err != nil {
		return 0, errors.New("CalculateFixedTokens failed")
	}
	val, _ := strconv.ParseUint(tokens, 10, 64)
	return val, nil
}
