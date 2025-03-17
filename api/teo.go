package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func runHttpServer(config *Config) {
	router := httprouter.New()
	api := ApiController{
		config: config,
	}

	api.AddRoutes(router)

	port := 8081
	if config.Api != nil && config.Api.Port != nil {
		port = *config.Api.Port
	}

	address := "0.0.0.0"
	if config.Api != nil && config.Api.Address != nil {
		address = *config.Api.Address
	}

	fullAddress := fmt.Sprintf("%s:%d", address, port)
	log.Printf("listening on address %s", fullAddress)

	err := http.ListenAndServe(fullAddress, router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args, err := parseArgs()

	if err != nil {
		log.Fatal(err)
	}

	config, err := parseConfig(args.Config)

	if err != nil {
		log.Fatal(err)
	}

	runHttpServer(config)
}
