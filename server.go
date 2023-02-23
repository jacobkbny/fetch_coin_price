package main

import (
	"log"
	"net/http"
	"time"
)

func server() {

	go func() {
		for {
			loc, _ := time.LoadLocation("Local")
			Now := time.Now()
			if Now == time.Date(Now.Year(), Now.Month(), Now.Day(), Now.Hour(), 00, 2, 0, loc) {
				if GetlimitBetting() == false {
					startBetting()
					time.Sleep(time.Minute * 58)
				} else {
					SetResult()
					time.Sleep(time.Minute * 58)
				}
			}
		}
	}()

	go func() {
		for {
			loc, _ := time.LoadLocation("Local")
			Now := time.Now()
			if Now == time.Date(Now.Year(), Now.Month(), Now.Day(), Now.Hour(), 5, 00, 0, loc) {
				EndBetting()
				time.Sleep(time.Minute * 58)
			}
		}
	}()

	hanlder()

	log.Fatal(http.ListenAndServe(":80", nil))
}

func hanlder() {
	http.HandleFunc("/getPrice", getPrice)
	// http.HandleFunc("/getEachBet", GetEachBet)
	// http.HandleFunc("/getTotalBet", GetTotalBet)
}
