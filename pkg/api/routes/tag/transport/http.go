package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mythio/go-rest-starter/pkg/api/routes/tag"
	"github.com/mythio/go-rest-starter/pkg/api/routes/tag/schema/req"
)

// HTTP represents tag http service
type HTTP struct {
	service tag.Service
}

// NewHTTP creates new tag http service
func NewHTTP(service tag.Service, router *gin.RouterGroup) {
	h := &HTTP{service}

	router.POST("/", h.create)
	router.GET("/search", h.search)
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

	tag, err := h.service.Create(tag.ReqCreate{
		Name:        reqBody.Name,
		Description: reqBody.Description,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tag": tag,
	})
}

func (h *HTTP) search(c *gin.Context) {
	tagName := c.Query("search")

	tags, err := h.service.Search(tagName)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tag": tags,
	})
}
