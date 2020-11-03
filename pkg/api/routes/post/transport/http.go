package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mythio/go-rest-starter/pkg/api/routes/post"
	"github.com/mythio/go-rest-starter/pkg/api/routes/post/schema/req"
)

// HTTP represents post http service
type HTTP struct {
	service post.Service
}

// NewHTTP creates new post http service
func NewHTTP(service post.Service, router *gin.RouterGroup) {
	h := &HTTP{service}

	router.POST("/", h.create)
}

func (h *HTTP) create(c *gin.Context) {
	reqBody := &req.Create{}
	if err := c.ShouldBindJSON(reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	userID := c.GetInt64("id")

	post, err := h.service.Create(post.ReqCreate{
		AuthorID: userID,
		Title:    reqBody.Title,
		Body:     reqBody.Body,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}
