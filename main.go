package main

import (
	_ "curiosity/rss"
	"fmt"
	"google.golang.org/appengine"
	"net/http"
	"time"
)

var lastExecuted = 1

func main() {
	// Tickers..
	ticker := time.NewTicker(50 * time.Minute)
	go func() {
		for t := range ticker.C {
			lastExecuted += 1
			fmt.Println("Tick at", t)
			fmt.Println("------------>", lastExecuted)
		}
	}()

	//URL Handlers
	http.HandleFunc("/", handleIndex)
	appengine.Main()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esaaa!!")
}
