package generic

import "github.com/thommil/animals-go-common/model"

// Configuration definition for generic providers
type Configuration struct {
	Alg    string
	Secret string
}

// Provider is the generic JWT Authentication mechanism
type Provider struct {
	Configuration *Configuration
}

// Authenticate implementation of Provider API
func (provider Provider) Authenticate(credentials interface{}) (*model.User, error) {
	return nil, nil
}
