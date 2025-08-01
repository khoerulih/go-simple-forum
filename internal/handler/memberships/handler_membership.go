package memberships

import "github.com/gin-gonic/gin"

type Handler struct {
	Engine *gin.Engine
}

func NewHandler(api *gin.Engine) *Handler {
	return &Handler{
		Engine: api,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Engine.Group("memberships")

	route.GET("ping", h.Ping)

}
