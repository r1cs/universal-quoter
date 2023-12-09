package uniswap_quoter

type V2Request struct {
	Amount  string `json:"amount"`
	Configs []struct {
		EnableFeeOnTransferFeeFetching bool     `json:"enableFeeOnTransferFeeFetching"`
		EnableUniversalRouter          bool     `json:"enableUniversalRouter"`
		Protocols                      []string `json:"protocols"`
		Recipient                      string   `json:"recipient"`
		RoutingType                    string   `json:"routingType"`
	} `json:"configs"`
	SendPortionEnabled bool   `json:"sendPortionEnabled"`
	TokenIn            string `json:"tokenIn"`
	TokenInChainID     int    `json:"tokenInChainId"`
	TokenOut           string `json:"tokenOut"`
	TokenOutChainID    int    `json:"tokenOutChainId"`
	Type               string `json:"type"`
}
