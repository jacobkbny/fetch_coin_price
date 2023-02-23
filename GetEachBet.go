package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEachBet(writer http.ResponseWriter, req *http.Request) {
	res, err := http.Get(NodeAddress + "/getEachBet")
	if err != nil {
		fmt.Println("Error GetEachBet: ", err)
	}
	var response map[string]map[string]string
	json.NewDecoder(res.Body).Decode(&response)
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		fmt.Println("error encoding:", err)
	}

}
