package keeper

import (
	"context"
	"math"
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/qlabs-xyz/qcore/x/pool/constants"
)

func (k Keeper) CalculateExponentialTokens(blockNumber int64) (string, error) {
	initialRate, err := sdkmath.LegacyNewDecFromStr(constants.InitialRate)
	if err != nil {
		return "", err
	}

	decay, err := sdkmath.LegacyNewDecFromStr(constants.Decay)
	if err != nil {
		return "", err
	}

	n := sdkmath.LegacyNewDec(blockNumber)
	decayN := decay.Mul(n)

	expArg := -decayN.MustFloat64()
	expResult := math.Exp(expArg)

	scaled := new(big.Float).SetFloat64(expResult)
	scaled.Mul(scaled, big.NewFloat(1e18))
	scaledInt := new(big.Int)
	scaled.Int(scaledInt)

	expVal := sdkmath.LegacyNewDecFromBigInt(scaledInt)
	expVal = expVal.QuoInt64(1e18)

	tokens := initialRate.Mul(expVal)

	return tokens.String(), nil
}

func (k Keeper) CalculateFixedTokens(ctx context.Context) (string, error) {

	// Fixed emission: (totalSupply * 0.02) / 365 / 17280
	apr, err := sdkmath.LegacyNewDecFromStr(constants.APR)
	if err != nil {
		return "", err
	}

	daysPerYear, err := sdkmath.LegacyNewDecFromStr(constants.DaysPerYear)
	if err != nil {
		return "", err
	}

	blocksPerDay, err := sdkmath.LegacyNewDecFromStr(constants.BlocksPerDay)
	if err != nil {
		return "", err
	}

	totalSupply, err := k.TotalSupply(ctx)
	if err != nil {
		return "", err
	}

	decTotalSupply := sdkmath.LegacyMustNewDecFromStr(totalSupply)
	tokens := decTotalSupply.Mul(apr).Quo(daysPerYear).Quo(blocksPerDay)

	return tokens.String(), nil
}
