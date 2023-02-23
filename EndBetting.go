package main

import (
	"fmt"
	"net/http"
)

func EndBetting() {
	_, err := http.Post(NodeAddress+"/endBetting", "text/plain", nil)
	if err != nil {
		fmt.Println("EndBetting Error:", err)
	}
}
