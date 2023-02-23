package main

import (
	"encoding/json"
	"net/http"
)

func getPrice(writer http.ResponseWriter, req *http.Request) {
	var request map[string]string
	json.NewDecoder(req.Body).Decode(&request)
	symbol := request["ticker"]
	if CoinPrices[symbol] == "0" {
		Price := FetchCoinPrice(symbol)
		json.NewEncoder(writer).Encode(Price)
	} else {
		json.NewEncoder(writer).Encode(CoinPrices[symbol])
	}
}
