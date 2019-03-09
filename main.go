package main

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"
)

func main(){
	http.HandleFunc("/", handleIndex)
	appengine.Main()
}

func handleIndex(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Esaaa!!")
}
