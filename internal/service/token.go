package service

func (s *Service) ValidateToken(token string) error {
	return s.repo.ValidateToken(token)
}
