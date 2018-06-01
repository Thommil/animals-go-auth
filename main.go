package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/thommil/animals-go-common/dao/mongo"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-auth/providers/facebook"
	"github.com/thommil/animals-go-auth/providers/google"
	"github.com/thommil/animals-go-auth/resources/authentication"
	"github.com/thommil/animals-go-common/config"
)

// Configuration definition for animals-go-auth
type Configuration struct {
	HTTP struct {
		Host string
		Port int
	}

	Mongo mongo.Configuration

	JWT authentication.JWTSettings

	Providers struct {
		Facebook facebook.Configuration
		Google   google.Configuration
	}
}

// Main of animals-go-ws
func main() {
	//Config
	configuration := &Configuration{}
	if err := config.LoadConfiguration("animals-go-auth", configuration); err != nil {
		log.Fatal(err)
	}

	//Mongo
	session, err := mongo.NewInstance(&configuration.Mongo)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//HTTP Server
	router := gin.Default()

	//Provider instances
	providers := map[string]authentication.Provider{
		"facebook": facebook.Provider{Configuration: &configuration.Providers.Facebook},
		"google":   google.Provider{Configuration: &configuration.Providers.Google},
	}

	//Resources
	authentication.New(router, providers, &configuration.JWT)

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
