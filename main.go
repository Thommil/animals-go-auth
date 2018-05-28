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
	"github.com/thommil/animals-go-common/config"
	"github.com/thommil/animals-go-common/model"
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

// Provider interface to define API of authentication providers
type Provider interface {
	Authenticate(token string) (*model.User, error)
}

// Main of animals-go-ws
func main() {
	//Config
	configuration := &Configuration{}
	if err := config.LoadConfiguration("animals-go-auth", configuration); err != nil {
		log.Fatal(err)
	}

	//Authentication instances
	providers := map[string]Provider{
		"generic":  generic.Provider{Configuration: &configuration.Providers.Generic},
		"facebook": facebook.Provider{Configuration: &configuration.Providers.Facebook},
		"google":   google.Provider{Configuration: &configuration.Providers.Google},
	}

	//HTTP Server
	router := gin.Default()

	//Middlewares
	user, err := providers["generic"].Authenticate("toto")
	fmt.Println(user.ID, err)

	//Start Server
	var serverAddress strings.Builder
	fmt.Fprintf(&serverAddress, "%s:%d", configuration.HTTP.Host, configuration.HTTP.Port)
	log.Fatal(endless.ListenAndServe(serverAddress.String(), router))
}
