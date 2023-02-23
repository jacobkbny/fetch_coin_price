package main

import (
	"bytes"
	"net/http"
)

func SetResult() {
	for k := range CoinPrices {
		Price := FetchCoinPrice(k)
		msg := bytes.NewBuffer([]byte(`{"price": ` + Price + `}`))
		http.Post(NodeAddress+"/setResult", "application/json", msg)
	}
}
