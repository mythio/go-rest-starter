package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mythio/go-rest-starter/pkg/api/routes/post"
	"github.com/mythio/go-rest-starter/pkg/api/routes/post/schema/req"
	"github.com/mythio/go-rest-starter/pkg/common/util/pagination"
)

// HTTP represents post http service
type HTTP struct {
	service post.Service
}

// NewHTTP creates new post http service
func NewHTTP(service post.Service, router *gin.RouterGroup) {
	h := &HTTP{service}

	router.POST("/", h.create)
	router.GET("/", h.getAll)
	router.GET("/:id", h.get)
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
		Tags:     reqBody.Tags,
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

func (h *HTTP) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	post, err := h.service.Get(id)
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

func (h *HTTP) getAll(c *gin.Context) {
	pageNo, err := strconv.ParseInt(c.Query("page_no"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	pageSize, err := strconv.ParseInt(c.Query("page_size"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	reqPage := pagination.ReqPagination{PageNo: int(pageNo), PageSize: int(pageSize)}

	posts, err := h.service.GetAll(reqPage)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
