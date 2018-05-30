package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/thommil/animals-go-auth/facebook"
	"github.com/thommil/animals-go-auth/google"
	"github.com/thommil/animals-go-auth/resources/authentication"
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

	//Provider instances
	providers := map[string]authentication.Provider{
		"facebook": facebook.Provider{Configuration: &configuration.Providers.Facebook},
		"google":   google.Provider{Configuration: &configuration.Providers.Google},
	}

	//Mongo
	session, err := mgo.Dial(configuration.Mongo.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//HTTP Server
	router := gin.Default()

	//Resources
	authentication.New(router, providers, session.DB(""), &configuration.JWT)

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
