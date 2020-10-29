package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mythio/go-rest-starter/pkg/api/auth"
	"github.com/mythio/go-rest-starter/pkg/common/model"
)

// HTTP represents user http service
type HTTP struct {
	service auth.Service
}

// NewHTTP creates new user http service
func NewHTTP(service auth.Service, e *gin.Engine) {
	h := &HTTP{service}
	router := e.Group("/auth")

	router.POST("/signup", h.signup)
	router.POST("/signin", h.signin)
}

func (h *HTTP) signup(c *gin.Context) {
	reqBody := &SignupReq{}
	if err := c.ShouldBindJSON(reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	if reqBody.Password != reqBody.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "passwords dont match",
		})

		return
	}

	user, err := h.service.Signup(model.User{
		Username:  reqBody.Username,
		Password:  reqBody.Password,
		Email:     reqBody.Email,
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *HTTP) signin(c *gin.Context) {
	reqBody := &SigninReq{}
	if err := c.ShouldBindJSON(reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	user, err := h.service.Signin(model.User{
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}