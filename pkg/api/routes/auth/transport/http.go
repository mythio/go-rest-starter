package transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mythio/go-rest-starter/pkg/api/middleware"
	"github.com/mythio/go-rest-starter/pkg/api/routes/auth"
	"github.com/mythio/go-rest-starter/pkg/api/routes/auth/schema/req"
)

// HTTP represents user http service
type HTTP struct {
	service auth.Service
}

// NewHTTP creates new user http service
func NewHTTP(service auth.Service, router *gin.RouterGroup) {
	h := &HTTP{service}

	router.POST("/signup", h.signup)
	router.POST("/signin", h.signin)

	router.Use(middleware.RequireToken())
	router.GET("/me", h.me)
}

func (h *HTTP) signup(c *gin.Context) {
	reqBody := &req.Signup{}
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

	authUser, err := h.service.Signup(auth.ReqSignup{
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

	c.JSON(http.StatusOK, authUser)
}

func (h *HTTP) signin(c *gin.Context) {
	reqBody := &req.Signin{}
	if err := c.ShouldBindJSON(reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	authUser, err := h.service.Signin(auth.ReqSignin{
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

	c.JSON(http.StatusOK, authUser)
}

func (h *HTTP) me(c *gin.Context) {
	meID := int64(c.GetInt("id"))
	user, err := h.service.Me(auth.ReqMe(meID))

	fmt.Print(meID)

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
