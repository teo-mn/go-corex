package utils

import "math/big"

// ToCrx number of CRX to Wei
func ToCrx(crx uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(crx), big.NewInt(1e18))
}
