package posts

import (
	"context"

	"github.com/khoerulih/go-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like to database")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comment to database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: *postDetail,
		LikeCount:  likeCount,
		Comments:   comments,
	}, nil
}
