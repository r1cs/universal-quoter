## install
```shell
go get -u github.com/r1cs/universal-quoter
```

```Go
// get quote route and call info of specific pair
uniswap_quoter.V2EndpointQuote(...)

// query eth price at eth mainnet, priced in USDT
universal_quoter.QueryEthPrice()
```