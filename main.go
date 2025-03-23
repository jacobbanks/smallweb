package main

import (
	"log"
	"net/http"
)


func main() {
	log.Print("Listening on port: 7777")
	log.Print("What are you interested in today?")
	handler := http.HandlerFunc(HelloServer)
	log.Fatal(http.ListenAndServe(":7777", handler))
}
