package universal_quoter

import (
	"testing"
)

func TestQueryEthPrice(t *testing.T) {
	price, err := QueryEthPrice()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(price)
}
