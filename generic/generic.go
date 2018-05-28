package generic

import (
	"github.com/dgrijalva/jwt-go"
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
	jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	return &user, nil
}
