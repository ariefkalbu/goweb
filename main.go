/*
Terminal : go mod init goweb
*/
package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {

	//ROUTING
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler) //as a root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/arief", handler.AriefHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Starting localhost:71")

	err := http.ListenAndServe(":71", mux)

	log.Fatal(err)
}
