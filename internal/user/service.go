package user

import (
	"avitoTechInternship/pkg/logging"
)

type Service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) Accrual(userDTO *UserDTO) error {
	err := s.repository.Accrual(*userDTO)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetBalance(userDTO *UserDTO) (User, error) {
	balance, err := s.repository.GetBalance(*userDTO)
	if err != nil {
		return User{}, err
	}
	return balance, nil
}
