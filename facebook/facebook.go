package facebook

import (
	"github.com/thommil/animals-go-common/model"
)

// Configuration definition for facebook providers
type Configuration struct{}

// Provider allows to check user entry against OAuth2 FB API
type Provider struct {
	Configuration *Configuration
}

// Authenticate implementation of Authentication API
func (provider Provider) Authenticate(credentials interface{}) (*model.User, error) {
	return nil, nil
}
