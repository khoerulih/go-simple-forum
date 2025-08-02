package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoerulih/go-simple-forum/internal/model/posts"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// get userid from auth (saved to context by middleware)
	userID := c.GetInt64("userID")

	if err := h.postSvc.CreatePost(ctx, userID, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
