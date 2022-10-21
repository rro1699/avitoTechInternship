package order

import "avitoTechInternship/pkg/logging"

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

func (s *Service) Reservation(dto *OrderDTO) error {
	err := s.repository.Reservation(*dto)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Recognition(dto *OrderDTO) error {
	err := s.repository.Recognition(*dto)
	if err != nil {
		return err
	}
	return nil
}
