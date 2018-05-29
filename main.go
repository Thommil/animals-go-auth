package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-auth/facebook"
	"github.com/thommil/animals-go-auth/generic"
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

	Providers struct {
		Generic  generic.Configuration
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

	//Authentication instances
	providers := map[string]authentication.Provider{
		"generic":  generic.Provider{Configuration: &configuration.Providers.Generic},
		"facebook": facebook.Provider{Configuration: &configuration.Providers.Facebook},
		"google":   google.Provider{Configuration: &configuration.Providers.Google},
	}

	//HTTP Server
	router := gin.Default()

	//Resources
	authentication := authentication.New(router, providers)
	authentication.Authenticate("test", "test")

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
