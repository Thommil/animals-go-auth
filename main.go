package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/config"
)

// Configuration definition for animals-go-auth
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
	if err := config.LoadConfiguration("animals-go-auth", configuration); err != nil {
		log.Fatal(err)
	}

	//HTTP Server
	router := gin.Default()

	//Middlewares

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
