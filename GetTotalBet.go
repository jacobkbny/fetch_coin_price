package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTotalBet(writer http.ResponseWriter, req *http.Request) {
	var request map[string]string
	json.NewDecoder(req.Body).Decode(&request)
	position := request["position"]
	msg := bytes.NewBuffer([]byte(position))
	res, err := http.Post(NodeAddress+"/getTotalbet", "application/json", msg)
	if err != nil {
		fmt.Println("Error GetTotalBet:", err)
	}
	var response map[string]string
	json.NewDecoder(res.Body).Decode(&response)
	Amount := response["amount"]
	json.NewEncoder(writer).Encode(Amount)

}
