package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type header struct {
	Authorization string `header:"Authorization"`
}

// TokenParser represents JWT token parser
type TokenParser interface {
	ParseToken(string) (*jwt.Token, error)
}

// CheckAuthToken makes JWT implement the CheckAuthToken interface.
func CheckAuthToken(t TokenParser) gin.HandlerFunc {
	return func(c *gin.Context) {
		var h header
		if err := c.ShouldBindHeader(&h); err != nil {
			c.Next()
			return
		}

		token, err := t.ParseToken(h.Authorization)
		if err != nil || !token.Valid {
			c.Next()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		id := int(claims["id"].(float64))
		email := claims["e"].(string)
		username := claims["u"].(string)

		c.Set("id", id)
		c.Set("email", email)
		c.Set("username", username)

		c.Next()
	}
}

// RequireToken enforces the need of 'Authorization' header in request
func RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get("id"); !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err": "no token passed",
			})
			return
		}
		if _, exists := c.Get("email"); !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err": "no token passed",
			})
			return
		}
		if _, exists := c.Get("username"); !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err": "no token passed",
			})
			return
		}
	}
}
