package universal_quoter

import "math/big"

func Power(decimal int) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
}

func Amount(decimal int, f float64) *big.Int {
	f = f * 1e6
	fe6 := new(big.Int).SetUint64(uint64(f))
	fe6.Mul(fe6, Power(decimal))
	out := fe6.Div(fe6, big.NewInt(1e6))
	return out
}

func FromRawAmount(decimal int, a *big.Int) float64 {
	//copy a
	aa := new(big.Int).Set(a)
	if aa.Sign() == -1 {
		aa.Mul(aa, big.NewInt(-1))
	}
	//a=a*1e6
	aa.Mul(aa, big.NewInt(1e6))
	//a=a/10^decimal
	aa.Div(aa, Power(decimal))
	fe6 := aa.Uint64()
	//a=a/1e6
	return float64(fe6) / 1e6
}
