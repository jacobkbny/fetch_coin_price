package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// 00시 00분이 되면 해야하는것
// 가격을 받아와서 그 가격을 nest.js 에게 보내줌
func CheckTime() {
	fmt.Println("CoinPrices length", len(CoinPrices))
	for k := range CoinPrices {
		FetchCoinPrice(k)
		Price := FetchCoinPrice(k)
		msg := bytes.NewBuffer([]byte(Price))
		http.Post("http://localhost:3000/SendPrice", "application/json", msg)
	}

}
