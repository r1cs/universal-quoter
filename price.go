package universal_quoter

import (
	"errors"
	uniswap_quoter "github.com/r1cs/universal-quoter/uniswap-quoter"
	"math/big"
)

func QueryEthPrice(amounts ...float64) (float64, error) {
	amount := 1.0
	if len(amounts) > 0 {
		amount = amounts[0]
	}
	ret, err := uniswap_quoter.V2EndpointQuote(1, "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266", "ETH", "0xdAC17F958D2ee523a2206206994597C13D831ec7", true, ETH(amount))
	if err != nil {
		return 0, err
	}
	out := new(big.Int)
	for _, route := range ret.Quote.Route {
		v, ok := new(big.Int).SetString(route[len(route)-1].AmountOut, 10)
		if !ok {
			return 0, errors.New("decode amountOut failed")
		}
		out.Add(out, v)
	}
	return FromRawAmount(6, out), nil
}

func ETH(f float64) *big.Int {
	return Amount(18, f)
}

func ToETH(a *big.Int) float64 {
	return FromRawAmount(18, a)
}

func Gwei(f float64) *big.Int {
	return Amount(9, f)
}

func ToGwei(a *big.Int) float64 {
	return FromRawAmount(9, a)
}
