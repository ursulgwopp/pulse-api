package service

import (
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (s *Service) NewPost(id int, req models.NewPostRequest) (models.Post, error) {
	for _, tag := range req.Tags {
		if !isValidTag(tag) {
			return models.Post{}, errors.ErrInvalidTag
		}
	}

	if len(req.Content) > 1000 {
		return models.Post{}, errors.ErrInvalidContent
	}

	login, err := s.repo.CheckLoginByUserId(id)
	if err != nil {
		return models.Post{}, errors.ErrInvalidContent
	}

	return s.repo.NewPost(login, req)
}
