//go:generate gojsonstruct -o generated_v2_resp.go -omitempty never -packagename uniswap_quoter -typename V2Result ./v2_resp.json
//go:generate gojsonstruct -o generated_v2_request.go -omitempty never -packagename uniswap_quoter -typename V2Request ./v2_request.json

package uniswap_quoter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
)

const V2_ENDPOINT = "https://api.uniswap.org/v2/quote"

type V2Config struct {
	Protocols                      []string `json:"protocols"`
	EnableUniversalRouter          bool     `json:"enableUniversalRouter"`
	RoutingType                    string   `json:"routingType"`
	Recipient                      string   `json:"recipient"`
	EnableFeeOnTransferFeeFetching bool     `json:"enableFeeOnTransferFeeFetching"`
}
type V2Param struct {
	TokenInChainId     uint64     `json:"tokenInChainId"`
	TokenOutChainId    uint64     `json:"tokenOutChainId"`
	TokenIn            string     `json:"tokenIn"`
	TokenOut           string     `json:"tokenOut"`
	Amount             string     `json:"amount"`
	SendPortionEnabled bool       `json:"sendPortionEnabled"`
	Type               string     `json:"type"`
	Configs            []V2Config `json:"configs"`
}

func genTypeStr(isExactIn bool) string {
	if isExactIn {
		return "EXACT_INPUT"
	}
	return "EXACT_OUTPUT"
}

func V2EndpointQuote(chainId int, recipient, tokenIn, tokenOut string, isExactIn bool, amount *big.Int) (ret V2Result, err error) {

	r := V2Param{
		TokenInChainId:     uint64(chainId),
		TokenOutChainId:    uint64(chainId),
		TokenOut:           tokenOut,
		TokenIn:            tokenIn,
		Amount:             amount.String(),
		SendPortionEnabled: true,
		Type:               genTypeStr(isExactIn),
		Configs: []V2Config{
			{
				Protocols:                      []string{"V2", "V3", "MIXED"},
				EnableUniversalRouter:          true,
				EnableFeeOnTransferFeeFetching: true,
				RoutingType:                    "CLASSIC",
				Recipient:                      recipient,
			},
		},
	}
	//r := V2Request{
	//	Amount:             amount.String(),
	//	SendPortionEnabled: true,
	//	TokenIn:            tokenIn,
	//	TokenOut:           tokenOut,
	//	TokenInChainID:     chainId,
	//	TokenOutChainID:    chainId,
	//	Type:               genTypeStr(isExactIn),
	//	Configs: []struct {
	//		EnableFeeOnTransferFeeFetching bool     `json:"enableFeeOnTransferFeeFetching"`
	//		EnableUniversalRouter          bool     `json:"enableUniversalRouter"`
	//		Protocols                      []string `json:"protocols"`
	//		Recipient                      string   `json:"recipient"`
	//		RoutingType                    string   `json:"routingType"`
	//	}{
	//		{
	//			EnableFeeOnTransferFeeFetching: true,
	//			EnableUniversalRouter:          true,
	//			Protocols:                      []string{"V2,V3,MIXED"},
	//			Recipient:                      recipient,
	//			RoutingType:                    "CLASSIC",
	//		},
	//	},
	//}

	data, err := json.Marshal(r)
	if err != nil {
		return ret, err
	}
	req, err := http.NewRequest("POST", V2_ENDPOINT, bytes.NewBuffer(data))
	if err != nil {
		return ret, err
	}
	req.Header.Set("Origin", "https://app.uniswap.org")
	req.Header.Set("Referer", "https://app.uniswap.org")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}
	if resp.StatusCode != http.StatusOK {
		return ret, fmt.Errorf("request failed: %s", b)
	}
	if err := json.Unmarshal(b, &ret); err != nil {
		return ret, err
	}

	if len(ret.AllQuotes) == 0 {
		return ret, errors.New("no quotes found")
	}
}
