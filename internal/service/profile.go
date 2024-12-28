package service

import (
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (s *Service) GetProfileByLogin(id int, login string) (models.UserProfile, error) {
	exists, err := s.repo.CheckLoginExists(login)
	if err != nil {
		return models.UserProfile{}, err
	}

	if !exists {
		return models.UserProfile{}, errors.ErrLoginDoesNotExist
	}

	userId, err := s.repo.CheckUserIdByLogin(login)
	if err != nil {
		return models.UserProfile{}, err
	}

	userProfile, err := s.repo.GetProfile(userId)
	if err != nil {
		return models.UserProfile{}, err
	}

	if id == userId {
		return userProfile, nil
	}

	if userProfile.IsPublic {
		return userProfile, nil
	}

	initialLogin, err := s.repo.CheckLoginByUserId(id)
	if err != nil {
		return models.UserProfile{}, err
	}

	friends, err := s.repo.ListFriends(userId, 1000000, 0)
	if err != nil {
		return models.UserProfile{}, err
	}

	if isFriend(friends, initialLogin) {
		return userProfile, err
	}

	return models.UserProfile{}, errors.ErrAccessDenied
}
