package main

import (
	"log"
	"net/http"

	"github.com/Adil-9/api.coingecko/hands"
)

func main() {

	http.HandleFunc("/", hands.HandleCoins)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Could not start server:", err.Error())
	}
}
