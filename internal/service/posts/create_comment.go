package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/khoerulih/go-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postID int64, userID int64, request posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.Itoa(int(userID)),
		UpdatedBy:      strconv.Itoa(int(userID)),
	}
	if err := s.postRepo.CreateComment(ctx, model); err != nil {
		log.Error().Err(err).Msg("failed to create comment to repository")
		return err
	}
	return nil
}
