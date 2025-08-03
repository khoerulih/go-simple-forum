package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/khoerulih/go-simple-forum/internal/middleware"
	"github.com/khoerulih/go-simple-forum/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID int64, userID int64, request posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID int64, userID int64, request posts.UserActivityRequest) error
	GetAllPost(ctx context.Context, pageSize int, pageIndex int) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error)
}

type Handler struct {
	Engine *gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Engine.Group("posts")

	route.Use(middleware.AuthMiddleware())
	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user-activity/:postID", h.UpsertUserActivity)
	route.GET("/", h.GetAllPost)
	route.GET("/:postID", h.GetPostByID)
}
