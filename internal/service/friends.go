package service

import (
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (s *Service) AddFriend(id int, login string) error {
	exists, err := s.repo.CheckLoginExists(login)
	if err != nil {
		return err
	}

	if !exists {
		return errors.ErrLoginDoesNotExist
	}

	return s.repo.AddFriend(id, login)
}

func (s *Service) RemoveFriend(id int, login string) error {
	exists, err := s.repo.CheckLoginExists(login)
	if err != nil {
		return err
	}

	if !exists {
		return errors.ErrLoginDoesNotExist
	}

	return s.repo.RemoveFriend(id, login)
}

func (s *Service) ListFriends(id int, limit int, offset int) ([]models.FriendInfo, error) {
	return s.repo.ListFriends(id, limit, offset)
}
