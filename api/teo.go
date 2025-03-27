package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr:    fullAddress,
		Handler: router,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handleGracefulShutdown(server)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	wg.Wait()
}

func handleGracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)

	<-stop

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server forced to shutdown: %v\n", err)
	}

	log.Println("server stopped")
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
