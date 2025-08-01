package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/khoerulih/go-simple-forum/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, error)
}

type Handler struct {
	Engine *gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Engine.Group("memberships")

	route.GET("ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)
}
