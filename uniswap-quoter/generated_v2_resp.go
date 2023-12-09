package uniswap_quoter

type V2Result struct {
	AllQuotes []struct {
		Quote struct {
			Amount                      string `json:"amount"`
			AmountDecimals              string `json:"amountDecimals"`
			BlockNumber                 string `json:"blockNumber"`
			GasPriceWei                 string `json:"gasPriceWei"`
			GasUseEstimate              string `json:"gasUseEstimate"`
			GasUseEstimateQuote         string `json:"gasUseEstimateQuote"`
			GasUseEstimateQuoteDecimals string `json:"gasUseEstimateQuoteDecimals"`
			GasUseEstimateUsd           string `json:"gasUseEstimateUSD"`
			HitsCachedRoutes            bool   `json:"hitsCachedRoutes"`
			MethodParameters            struct {
				Calldata string `json:"calldata"`
				To       string `json:"to"`
				Value    string `json:"value"`
			} `json:"methodParameters"`
			PortionAmount                      string `json:"portionAmount"`
			PortionAmountDecimals              string `json:"portionAmountDecimals"`
			PortionBips                        int    `json:"portionBips"`
			Quote                              string `json:"quote"`
			QuoteDecimals                      string `json:"quoteDecimals"`
			QuoteGasAdjusted                   string `json:"quoteGasAdjusted"`
			QuoteGasAdjustedDecimals           string `json:"quoteGasAdjustedDecimals"`
			QuoteGasAndPortionAdjusted         string `json:"quoteGasAndPortionAdjusted"`
			QuoteGasAndPortionAdjustedDecimals string `json:"quoteGasAndPortionAdjustedDecimals"`
			QuoteID                            string `json:"quoteId"`
			RequestID                          string `json:"requestId"`
			Route                              [][]struct {
				Address      string `json:"address"`
				AmountIn     string `json:"amountIn"`
				AmountOut    string `json:"amountOut"`
				Fee          string `json:"fee"`
				Liquidity    string `json:"liquidity"`
				SqrtRatioX96 string `json:"sqrtRatioX96"`
				TickCurrent  string `json:"tickCurrent"`
				TokenIn      struct {
					Address  string `json:"address"`
					ChainID  int    `json:"chainId"`
					Decimals string `json:"decimals"`
					Symbol   string `json:"symbol"`
				} `json:"tokenIn"`
				TokenOut struct {
					Address  string `json:"address"`
					ChainID  int    `json:"chainId"`
					Decimals string `json:"decimals"`
					Symbol   string `json:"symbol"`
				} `json:"tokenOut"`
				Type string `json:"type"`
			} `json:"route"`
			RouteString      string  `json:"routeString"`
			SimulationError  bool    `json:"simulationError"`
			SimulationStatus string  `json:"simulationStatus"`
			Slippage         float64 `json:"slippage"`
			TradeType        string  `json:"tradeType"`
		} `json:"quote"`
		Routing string `json:"routing"`
	} `json:"allQuotes"`
	Quote struct {
		Amount                      string `json:"amount"`
		AmountDecimals              string `json:"amountDecimals"`
		BlockNumber                 string `json:"blockNumber"`
		GasPriceWei                 string `json:"gasPriceWei"`
		GasUseEstimate              string `json:"gasUseEstimate"`
		GasUseEstimateQuote         string `json:"gasUseEstimateQuote"`
		GasUseEstimateQuoteDecimals string `json:"gasUseEstimateQuoteDecimals"`
		GasUseEstimateUsd           string `json:"gasUseEstimateUSD"`
		HitsCachedRoutes            bool   `json:"hitsCachedRoutes"`
		MethodParameters            struct {
			Calldata string `json:"calldata"`
			To       string `json:"to"`
			Value    string `json:"value"`
		} `json:"methodParameters"`
		PortionAmount                      string `json:"portionAmount"`
		PortionAmountDecimals              string `json:"portionAmountDecimals"`
		PortionBips                        int    `json:"portionBips"`
		Quote                              string `json:"quote"`
		QuoteDecimals                      string `json:"quoteDecimals"`
		QuoteGasAdjusted                   string `json:"quoteGasAdjusted"`
		QuoteGasAdjustedDecimals           string `json:"quoteGasAdjustedDecimals"`
		QuoteGasAndPortionAdjusted         string `json:"quoteGasAndPortionAdjusted"`
		QuoteGasAndPortionAdjustedDecimals string `json:"quoteGasAndPortionAdjustedDecimals"`
		QuoteID                            string `json:"quoteId"`
		RequestID                          string `json:"requestId"`
		Route                              [][]struct {
			Address      string `json:"address"`
			AmountIn     string `json:"amountIn"`
			AmountOut    string `json:"amountOut"`
			Fee          string `json:"fee"`
			Liquidity    string `json:"liquidity"`
			SqrtRatioX96 string `json:"sqrtRatioX96"`
			TickCurrent  string `json:"tickCurrent"`
			TokenIn      struct {
				Address  string `json:"address"`
				ChainID  int    `json:"chainId"`
				Decimals string `json:"decimals"`
				Symbol   string `json:"symbol"`
			} `json:"tokenIn"`
			TokenOut struct {
				Address  string `json:"address"`
				ChainID  int    `json:"chainId"`
				Decimals string `json:"decimals"`
				Symbol   string `json:"symbol"`
			} `json:"tokenOut"`
			Type string `json:"type"`
		} `json:"route"`
		RouteString      string  `json:"routeString"`
		SimulationError  bool    `json:"simulationError"`
		SimulationStatus string  `json:"simulationStatus"`
		Slippage         float64 `json:"slippage"`
		TradeType        string  `json:"tradeType"`
	} `json:"quote"`
	RequestID string `json:"requestId"`
	Routing   string `json:"routing"`
}
