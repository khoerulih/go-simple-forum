package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoerulih/go-simple-forum/internal/model/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	err := h.membershipSvc.SignUp(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, nil)
}
