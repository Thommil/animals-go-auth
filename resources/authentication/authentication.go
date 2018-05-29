// Package authentication defines authentication routing
package authentication

import (
	"github.com/thommil/animals-go-common/api"
	"github.com/thommil/animals-go-common/model"

	"github.com/gin-gonic/gin"
)

// Provider interface defines API for an authentication provider
type Provider interface {
	Authenticate(credentials interface{}) (*model.User, error)
}

type authentication struct {
	group     *gin.RouterGroup
	providers map[string]Provider
}

// New creates new Routable implementation for authentication features
func New(engine *gin.Engine, providers map[string]Provider) resource.Routable {
	authentication := &authentication{group: engine.Group("/"), providers: providers}
	{
		authentication.group.POST("/public/authenticate", authentication.publicAuthenticate)
		authentication.group.GET("/private/authenticate", authentication.privateAuthenticate)
	}
	return authentication
}

// GetGroup implementation of resource.Routable
func (authentication *authentication) GetGroup() *gin.RouterGroup {
	return authentication.group
}

func (authentication *authentication) publicAuthenticate(c *gin.Context) {
	//Check provider
	// providerImpl, ok := authentication.providers[provider]
	// if !ok {
	// 	return nil, fmt.Errorf("provider '%s' not found", provider)
	// }
	// return providerImpl.Authenticate(token)
}

func (authentication *authentication) privateAuthenticate(c *gin.Context) {
	//Check token
}
