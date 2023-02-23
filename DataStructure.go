package main

type Kline struct {
	OpenTime                 float64
	Open                     string
	High                     string
	Low                      string
	Close                    string
	Volume                   string
	CloseTime                float64
	QuoteAssetVolume         string
	NumberOfTrades           float64
	TakerBuyBaseAssetVolume  string
	TakerBuyQuoteAssetVolume string
}

type Request struct {
	ticker string
}

type Message struct {
	price string
}

var CoinPrices map[string]string

var NodeAddress = "http://localhost:3000"
