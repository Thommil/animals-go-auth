package google

import "github.com/thommil/animals-go-common/model"

// Configuration definition for facebook providers
type Configuration struct{}

// Provider allows to check user entry against Google JWT token
type Provider struct {
	Configuration *Configuration
}

// Authenticate implementation of Provider API
func (provider Provider) Authenticate(credentials interface{}) (*model.User, error) {
	return nil, nil
}
