package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thommil/animals-go-common/api"
	"github.com/thommil/animals-go-common/model"
)

// Provider interface to define API of authentication providers
type Provider interface {
	Authenticate(token string) (*model.User, error)
}

// Authentication handler
type Authentication struct {
	Providers *map[string]Provider
	Resource  *api.Resource
}

// ApplyRoutes implements IRoutable interface
func (authentication *Authentication) ApplyRoutes() *api.Resource {
	authentication.Resource.Engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "AUTH OK")
	})
	return authentication.Resource
}
