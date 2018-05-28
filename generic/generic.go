package generic

import (
	"github.com/thommil/animals-go-common/model"
)

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
func (provider Provider) Authenticate(token string) (*model.User, error) {
	user := model.User{ID: "titi"}
	return &user, nil
}
