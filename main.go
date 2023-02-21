package main

func init() {
	CoinPrices = make(map[string]string)
	CoinPrices["BTC"] = "0"
}

func main() {
	server()
}
