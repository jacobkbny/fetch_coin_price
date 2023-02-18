package main

import (
	"log"
	"net/http"
)

func server() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hanlder() {
	http.HandleFunc("/getPrice", getPrice)
}
