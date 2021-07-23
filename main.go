package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HealthHandler)
	http.HandleFunc("/webhook", WebhookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
