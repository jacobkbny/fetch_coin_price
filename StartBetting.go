package main

import (
	"bytes"
	"net/http"
)

// 00시 00분이 되면 해야하는것
// 가격을 받아와서 그 가격을 Node 에게 보내줌
func startBetting() {
	// if getlimitBetting is true then
	for k := range CoinPrices {
		Price := FetchCoinPrice(k)
		msg := bytes.NewBuffer([]byte(`{"price": ` + Price + `}`))
		http.Post(NodeAddress+"/startBetting", "application/json", msg)
	}
}
