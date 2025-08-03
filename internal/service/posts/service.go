package posts

import (
	"context"

	"github.com/khoerulih/go-simple-forum/internal/configs"
	"github.com/khoerulih/go-simple-forum/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	GetAllPost(ctx context.Context, limit int, offset int) (posts.GetAllPostResponse, error)

	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error

	GetPostByID(ctx context.Context, id int64) (*posts.Post, error)
	CountLikeByPostID(ctx context.Context, postID int64) (int, error)
	GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comment, error)
}

type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
