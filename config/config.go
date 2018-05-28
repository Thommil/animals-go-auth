package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
	"path/filepath"
	"strings"
)

// Configuration definition for animals-go-ws
type Configuration struct {
	Http struct {
		Host string
		Port int
	}

	Mongo struct {
		Url string
	}
}

// Load configuration from json file to Configuration
func LoadConfiguration() (Configuration, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	var filePath strings.Builder
	fmt.Fprintf(&filePath, "%s/config/animals-go-auth.json", exPath)
	configuration := Configuration{}
	err = gonfig.GetConf(filePath.String(), &configuration)
	return configuration, err
}
