package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-auth/config"
	"log"
	"strings"
)

// Main of animals-go-auth
func main() {
	//Config
	config, err := config.LoadConfiguration()

	if err != nil {
		log.Fatal(err)
	}

	//HTTP Server
	router := gin.Default()

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", config.Http.Host, config.Http.Port)
	log.Printf("Starting HTTP server on %s\n", serverAddress.String())
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
