package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/webhook", webhookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
