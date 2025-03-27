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

	fullAddress := fmt.Sprintf("%s:%d", *config.Api.Address, *config.Api.Port)
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
