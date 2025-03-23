package main

import (
	"net/http"
	"fmt"
	"log"
)
//
// func ListenAndServer(addr: string, handler: Handler) error
//
//
// type Handler interface {
// 	serveHTTP(ResponseWriter, *Request)
// }

func HelloServer(writer http.ResponseWriter, req *http.Request) {
	log.Print("Request Recieved")
	fmt.Fprint(writer, "Hello World")
}


