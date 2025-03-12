package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config, err := parseConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	api := ApiController{
		config: config,
	}

	api.AddRoutes(router)

	port := 8080
	if config.Http != nil && config.Http.Port != nil {
		port = *config.Http.Port
	}

	address := "0.0.0.0"
	if config.Http != nil && config.Http.Address != nil {
		address = *config.Http.Address
	}

	fullAddress := fmt.Sprintf("%s:%d", address, port)
	log.Printf("listening on address %s", fullAddress)

	err = http.ListenAndServe(fullAddress, router)
	if err != nil {
		log.Fatal(err)
	}
}
