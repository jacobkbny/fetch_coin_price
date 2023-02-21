package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func FetchCoinPrice(ticker string) string {
	symbol := strings.ToUpper(ticker)
	symbol_To_binance := symbol + "USDT"
	interval := "1s"
	startTime := getStartTime()
	endTime := getEndTime()

	url := fmt.Sprintf("https://api.binance.com/api/v3/klines?symbol=%s&interval=%s&startTime=%s&endTime=%s", symbol_To_binance, interval, startTime, endTime)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer response.Body.Close()
	var klines [][]interface{}
	err = json.NewDecoder(response.Body).Decode(&klines)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, klineData := range klines {
		kline := Kline{
			OpenTime:                 klineData[0].(float64),
			Open:                     klineData[1].(string),
			High:                     klineData[2].(string),
			Low:                      klineData[3].(string),
			Close:                    klineData[4].(string),
			Volume:                   klineData[5].(string),
			CloseTime:                klineData[6].(float64),
			QuoteAssetVolume:         klineData[7].(string),
			NumberOfTrades:           klineData[8].(float64),
			TakerBuyBaseAssetVolume:  klineData[9].(string),
			TakerBuyQuoteAssetVolume: klineData[10].(string),
		}
		CoinPrices[symbol] = kline.Open
		break
	}
	return CoinPrices[symbol]
}
