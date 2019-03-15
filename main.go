package main

import (
	"curiosity/rss"
	"fmt"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	// Start beating...
	fmt.Println("Helloooo...!")
	rss.StartBeat()

	//URL Handlers
	http.HandleFunc("/", handleIndex)
	appengine.Main()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esaaa!!")
}
