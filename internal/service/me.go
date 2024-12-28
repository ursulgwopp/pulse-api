package service

import "github.com/ursulgwopp/pulse-api/internal/models"

func (s *Service) GetProfile(id int) (models.UserProfile, error) {
	return s.repo.GetProfile(id)
}

func (s *Service) UpdateProfile(id int, req models.UpdateProfileRequest) (models.UserProfile, error) {
	if req.CountryCode != nil {
		if err := validateCountryCode(s, *req.CountryCode); err != nil {
			return models.UserProfile{}, err
		}
	}

	if req.Phone != nil {
		if err := validatePhone(s, *req.Phone); err != nil {
			return models.UserProfile{}, err
		}
	}

	if req.Image != nil {
		if err := validateImage(*req.Image); err != nil {
			return models.UserProfile{}, err
		}
	}

	return s.repo.UpdateProfile(id, req)
}

func (s *Service) UpdatePassword(id int, req models.UpdatePasswordRequest) error {
	if err := validatePassword(req.NewPassword); err != nil {
		return err
	}

	req.OldPassword = generatePasswordHash(req.OldPassword)
	req.NewPassword = generatePasswordHash(req.NewPassword)

	if err := s.repo.UpdatePassword(id, req); err != nil {
		return err
	}

	if err := s.repo.KillTokens(id); err != nil {
		return err
	}

	return nil
}
