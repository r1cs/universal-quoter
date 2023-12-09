package uniswap_quoter

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func TestV2Oracle(t *testing.T) {
	ret, err := V2EndpointQuote(5, "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266", "0xdD69DB25F6D620A7baD3023c5d32761D353D3De9", "0xa9A4F7f968A7CBD2d602487F42D2514713FF7556", true, big.NewInt(params.Ether))
	if err != nil {
		t.Fatal(err)
	}

	out := big.NewInt(0)
	for _, route := range ret.AllQuotes[0].Quote.Route {
		last := route[len(route)-1]
		t.Log(last.TokenOut.Symbol)
		i, ok := new(big.Int).SetString(last.AmountOut, 10)
		if !ok {
			t.Fatal(1)
		}
		out.Add(out, i)
	}

	t.Log(out.String())
	if len(ret.AllQuotes) == 0 {
		t.Fatal(1)
	}
	b, _ := json.MarshalIndent(ret, "", " ")
	t.Log(string(b))

}
