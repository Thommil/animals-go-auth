// Package authentication defines authentication routing
package authentication

import (
	"fmt"

	"github.com/thommil/animals-go-common/api"

	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/model"
)

// Provider interface defines API for an authentication provider
type Provider interface {
	Authenticate(token string) (*model.User, error)
}

type authentication struct {
	group     *gin.RouterGroup
	providers map[string]Provider
}

// New creates new Routable implementation for authentication features
func New(engine *gin.Engine, providers map[string]Provider) resource.Routable {
	authentication := &authentication{group: engine.Group("/"), providers: providers}
	{

	}
	return authentication
}

// GetGroup implementation of resource.Routable
func (authentication *authentication) GetGroup() *gin.RouterGroup {
	return authentication.group
}

func (authentication *authentication) authenticate(provider string, token string) (*model.User, error) {
	providerImpl, ok := authentication.providers[provider]
	if !ok {
		return nil, fmt.Errorf("provider '%s' not found", provider)
	}
	return providerImpl.Authenticate(token)
}
