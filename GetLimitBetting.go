package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetlimitBetting() bool {
	res, err := http.Get(NodeAddress + "/getlimitBetting")
	if err != nil {
		fmt.Println("Error getting limit:", err)
	}
	var response map[string]bool
	json.NewDecoder(res.Body).Decode(&response)
	result := response["limitBetting"]
	fmt.Println("result: ", result)
	return result

}
