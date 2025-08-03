package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/khoerulih/go-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID int64, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.Itoa(int(userID)),
		UpdatedBy: strconv.Itoa(int(userID)),
	}
	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user activity from database")
		return err
	}

	if userActivity == nil {
		// create user activity
		if !request.IsLiked {
			return errors.New("postingan ini belum pernah anda like sebelumnya")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		// update user activity
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("error create or update user activity to database")
		return err
	}
	return nil
}
