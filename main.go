package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/disturb16/graphql_golang/internal/handlers"
	"github.com/disturb16/graphql_golang/internal/services"
	"github.com/disturb16/graphql_golang/settings"
)

func main() {
	//Get project configurations
	config, err := settings.Configuration("")
	if err != nil {
		log.Fatal(err)
	}

	// initialize service
	service, err := services.New("mysql", config)
	if err != nil {
		log.Fatalf("Could not start service... %v", err)
	}

	//Create main handler
	handler := handlers.New(service)

	//Start server
	portToListen := strconv.Itoa(config.Port)
	fmt.Println("server running on port " + portToListen)
	log.Fatal(http.ListenAndServe(":"+portToListen, handler.Router()))
}
