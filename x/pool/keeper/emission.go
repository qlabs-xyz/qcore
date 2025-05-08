package keeper

import (
	"context"
	"fmt"
	"math"
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/qlabs-xyz/qcore/x/pool/constants"
)

func (k Keeper) CalculateExponentialBlockEmission(blockNumber int64) (string, error) {
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

func (k Keeper) CalculateFixedBlockEmission(ctx context.Context) (string, error) {

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

	fmt.Println("555555555555555", totalSupply)
	fmt.Println("6666666666666666666666", err)

	if totalSupply <= "0" {
		return "", fmt.Errorf("total supply must be greater than zero")
	}

	totalSupply = ""

	decTotalSupply := sdkmath.LegacyMustNewDecFromStr(totalSupply)
	emissionPerBlock := decTotalSupply.Mul(apr).Quo(daysPerYear).Quo(blocksPerDay)

	fmt.Println("00000000000000000000000000")
	fmt.Println("1111111111111111111111111", emissionPerBlock)
	fmt.Println("22222222222222222222222", emissionPerBlock.String())

	return emissionPerBlock.String(), nil
}

func (k Keeper) CalculateFixedDailyEmission(ctx context.Context) (string, error) {

	// Fixed emission: (totalSupply * 0.02) / 365
	apr, err := sdkmath.LegacyNewDecFromStr(constants.APR)
	if err != nil {
		return "", err
	}

	daysPerYear, err := sdkmath.LegacyNewDecFromStr(constants.DaysPerYear)
	if err != nil {
		return "", err
	}

	totalSupply, err := k.TotalSupply(ctx)
	if err != nil {
		return "", err
	}

	if totalSupply <= "0" {
		return "", fmt.Errorf("total supply must be greater than zero")
	}

	decTotalSupply := sdkmath.LegacyMustNewDecFromStr(totalSupply)
	dailyEmissionPerBlock := decTotalSupply.Mul(apr).Quo(daysPerYear)

	return dailyEmissionPerBlock.String(), nil
}

func (k Keeper) CalculateFixedAnnualEmission(ctx context.Context) (string, error) {

	// Fixed emission: (totalSupply * 0.02)
	apr, err := sdkmath.LegacyNewDecFromStr(constants.APR)
	if err != nil {
		return "", err
	}

	totalSupply, err := k.TotalSupply(ctx)
	if err != nil {
		return "", err
	}

	decTotalSupply := sdkmath.LegacyMustNewDecFromStr(totalSupply)
	annualEmissionPerBlock := decTotalSupply.Mul(apr)

	return annualEmissionPerBlock.String(), nil
}
