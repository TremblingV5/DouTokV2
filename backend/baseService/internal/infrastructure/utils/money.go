package utils

import "math/big"

var (
	LevelYuan2Li = big.NewFloat(1000)
)

func Ui642Money(amount uint64) *big.Float {
	return new(big.Float).Quo(big.NewFloat(0).SetUint64(amount), LevelYuan2Li)
}

func Money2Ui64(amount *big.Float) uint64 {
	result, _ := new(big.Float).Mul(amount, LevelYuan2Li).Uint64()
	return result
}
