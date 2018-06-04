// Package authentication defines authentication routing
package authentication

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/thommil/animals-go-common/api"
	"github.com/thommil/animals-go-common/model"

	"github.com/gin-gonic/gin"
)

// Provider interface defines API for an authentication provider
type Provider interface {
	Authenticate(credentials interface{}) (*model.User, error)
}

// JWTSettings defines JWT configuration
type JWTSettings struct {
	Secret  string        `json:"secret"`
	Expired time.Duration `json:"expired"`
	Issuer  string        `json:"issuer"`
}

// JWT token used bu local authentication (private)
type JWT struct {
	Token     string
	ExpiresAt int64
}

type authentication struct {
	group       *gin.RouterGroup
	providers   map[string]Provider
	jwtSettings *JWTSettings
}

// New creates new Routable implementation for authentication features
func New(engine *gin.Engine, providers map[string]Provider, jwtSettings *JWTSettings) resource.Routable {
	authentication := &authentication{group: engine.Group("/"), providers: providers, jwtSettings: jwtSettings}
	{
		authentication.group.GET("/public/authenticate/:provider/:tokenString", authentication.publicAuthenticate)
		authentication.group.GET("/private/authenticate/:tokenString", authentication.privateAuthenticate)
	}
	return authentication
}

// GetGroup implementation of resource.Routable
func (authentication *authentication) GetGroup() *gin.RouterGroup {
	return authentication.group
}

func (authentication *authentication) publicAuthenticate(c *gin.Context) {
	provider := c.Param("provider")
	tokenString := c.Param("tokenString")
	if user, err := authentication.providers[provider].Authenticate(tokenString); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
		})
	} else {
		claims := &jwt.StandardClaims{
			Subject:   user.ID.Hex(),
			Issuer:    authentication.jwtSettings.Issuer,
			ExpiresAt: time.Now().Add(time.Second * authentication.jwtSettings.Expired).Unix(),
			IssuedAt:  time.Now().Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if ss, err := token.SignedString([]byte(authentication.jwtSettings.Secret)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, JWT{Token: ss, ExpiresAt: claims.ExpiresAt})
		}
	}
}

func (authentication *authentication) privateAuthenticate(c *gin.Context) {
	tokenString := c.Param("tokenString")
	// Check token headers
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(authentication.jwtSettings.Secret), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
		})
	} else {
		//Check token validity
		if token.Valid {
			//Check token claims
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				//Check user ID validity
				if userID, ok := claims["sub"]; ok {
					if user, err := model.FindUserByID(userID.(string)); err != nil {
						c.JSON(http.StatusUnauthorized, gin.H{
							"code":    http.StatusUnauthorized,
							"message": err.Error(),
						})
					} else {
						c.JSON(http.StatusOK, user)
					}
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{
						"code":    http.StatusUnauthorized,
						"message": "Missing 'sub' claim",
					})
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": err.Error(),
				})
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "Token is invalid",
				})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "Token is expired",
				})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": err.Error(),
				})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
		}
	}
}
