package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	//"github.com/thommil/animals-go-auth/api"
	"github.com/thommil/animals-go-common/config"
)

// Configuration definition for animals-go-ws
type Configuration struct {
	HTTP struct {
		Host string
		Port int
	}

	Mongo struct {
		URL string
	}
}

// Main of animals-go-ws
func main() {
	//Config
	configuration := &Configuration{}
	err := config.LoadConfiguration("animals-go-auth", configuration)

	if err != nil {
		log.Fatal(err)
	}

	//HTTP Server
	router := gin.Default()

	//Middlewares

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Printf("Starting HTTP server on %s\n", serverAddress.String())
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
