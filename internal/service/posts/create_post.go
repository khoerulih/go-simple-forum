package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/khoerulih/go-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHashtags, ",")

	now := time.Now()

	model := posts.PostModel{
		UserID:       userID,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postHashtags,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    strconv.Itoa(int(userID)),
		UpdatedBy:    strconv.Itoa(int(userID)),
	}

	if err := s.postRepo.CreatePost(ctx, model); err != nil {
		log.Error().Err(err).Msg("error create post to a repository")
		return err
	}

	return nil
}
