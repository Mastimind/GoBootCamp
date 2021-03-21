package main

import (
	"log"
	"net/http"
)

func main() {
	// in cmd run  " curl -v localhost:6060/"  or in browser localhost:6060/
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Vanakam Ullagam !!")
	})
	// in cmd run  " curl -v localhost:6060/vidia"  or in browser localhost:6060/vidai

	http.HandleFunc("/vidai", func(http.ResponseWriter, *http.Request) {
		log.Println("Nall Vidai Ullagam !!")
	})

	http.ListenAndServe(":6060", nil)
}
