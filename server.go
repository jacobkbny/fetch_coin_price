package main

import (
	"log"
	"net/http"
)

func server() {

	// go func() {
	// 	for {
	// 		CheckTime()
	// 		time.Sleep(time.Second * 10)
	// 	}
	// }()

	hanlder()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hanlder() {
	http.HandleFunc("/getPrice", getPrice)
}
