package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/disturb16/graphql_golang/internal/handlers"
	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/disturb16/graphql_golang/settings"
	"github.com/disturb16/graphql_golang/schema"
)

func main() {
	//Get project configurations
	config, err := settings.Configuration("")
	if err != nil {
		log.Fatal(err)
	}

	var retries = 5
	var service *services.Service
	var waitTime = time.Second * 5

	// initialize service
	// if an error occurs, will try to reconnect
	for retries > 0 {
		retries--
		service, err = services.New("mysql", config)

		if err != nil {
			if retries == 0 {
				log.Fatalf("Could not start service... %v", err)
			}

			log.Printf("Could not start service... %v", err)
			log.Printf("Retrying in %v ...", waitTime)
			time.Sleep(waitTime)
		} else {
			break
		}
	}

	graphqlHandler, err := schema.NewHandler()

	if err != nil {
		log.Fatal(err)
	}

	//Create main handler
	handler := handlers.New(service, graphqlHandler)

	//Start server
	portToListen := strconv.Itoa(config.Port)
	fmt.Println("server running on port " + portToListen)
	log.Fatal(http.ListenAndServe(":"+portToListen, handler.Router()))
}
