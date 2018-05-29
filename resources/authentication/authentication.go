package authentication

import (
	"github.com/thommil/animals-go-common/api"

	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/model"
)

// Provider interface to define API of authentication providers
type Provider interface {
	Authenticate(token string) (*model.User, error)
}

type Authentication interface {
	resource.Routable
	Authenticate(provider string, token string) (*model.User, error)
}

type authentication struct {
	group     *gin.RouterGroup
	providers map[string]Provider
}

// ApplyRoutes implements IRoutable interface
func New(engine *gin.Engine, providers map[string]Provider) Authentication {
	authentication := &authentication{group: engine.Group("/"), providers: providers}
	{

	}
	return authentication
}

// GetGroup implementation of IRoutable
func (authentication *authentication) GetGroup() *gin.RouterGroup {
	return authentication.group
}

// GetGroup implementation of IRoutable
func (authentication *authentication) Authenticate(provider string, token string) (*model.User, error) {
	return authentication.providers[provider].Authenticate(token)
}
