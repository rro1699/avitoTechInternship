package user

import (
	"avitoTechInternship/internal/order"
	"avitoTechInternship/pkg/logging"
	"fmt"
	"github.com/shopspring/decimal"
)

type Service struct {
	repositoryOrder order.Repository
	repositoryUser  Repository
	logger          *logging.Logger
}

func NewService(repositoryOrder order.Repository, repositoryUser Repository, logger *logging.Logger) *Service {
	return &Service{
		repositoryOrder: repositoryOrder,
		repositoryUser:  repositoryUser,
		logger:          logger,
	}
}

func (s *Service) Accrual(dto *UserDTO) (string, error) {
	s.logger.Debug("Accrual money")
	localUser := s.repositoryUser.GetUserById(*dto)

	if (User{} != localUser) {

		curBalanceDecimal, _ := decimal.NewFromString(localUser.CurB)
		coastDecimal, _ := decimal.NewFromString(dto.Coast)
		curBalanceDecimal = curBalanceDecimal.Add(coastDecimal)

		localUser.CurB = curBalanceDecimal.String()

		err := s.repositoryUser.UpdateUser(localUser)
		if err != nil {
			return "", err
		}
		return fmt.Sprint(localUser.Id), nil
	} else {
		s.logger.Debug("Insert new user")
		curBalanceDecimal, _ := decimal.NewFromString(dto.Coast)

		var newUser User
		newUser.CurB = curBalanceDecimal.String()
		newUser.ResB = decimal.Zero.String()

		userId, err := s.repositoryUser.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		return userId, nil
	}
}

func (s *Service) GetBalance(dto *UserDTO) (User, error) {
	s.logger.Debug("Get balance")
	localUser := s.repositoryUser.GetUserById(*dto)
	if (User{} != localUser) {
		return localUser, nil
	} else {
		return User{}, fmt.Errorf("user not found")
	}
}

func (s *Service) Reservation(dto *order.OrderDTO) (string, error) {
	s.logger.Debug("Reservation money")

	localUser := s.repositoryUser.GetUserById(UserDTO{Id: dto.IdUser})
	if (User{} != localUser) {
		curBalanceDecimal, _ := decimal.NewFromString(localUser.CurB)
		resBalanceDecimal, _ := decimal.NewFromString(localUser.ResB)
		coastDecimal, _ := decimal.NewFromString(dto.Coast)

		if sub := curBalanceDecimal.Sub(coastDecimal); sub.Cmp(decimal.Zero) < 0 {
			return "", fmt.Errorf("user don't have enough money on balance")
		} else {
			resBalanceDecimal = resBalanceDecimal.Add(coastDecimal)
			curBalanceDecimal = curBalanceDecimal.Sub(coastDecimal)
			localUser.CurB = curBalanceDecimal.String()
			localUser.ResB = resBalanceDecimal.String()

			err := s.repositoryUser.UpdateUser(localUser)
			if err != nil {
				return "", err
			}

			err = s.repositoryOrder.Create(*dto)
			if err != nil {
				return "", err
			}

			return "Reservation ended successfully", nil
		}
	} else {
		return "", fmt.Errorf("user not found")
	}
}

func (s *Service) Recognition(dto *order.OrderDTO) (string, error) {
	s.logger.Debug("Recognition order")

	localUser := s.repositoryUser.GetUserById(UserDTO{Id: dto.IdUser})
	if (User{} != localUser) {
		resBalanceDecimal, _ := decimal.NewFromString(localUser.ResB)
		coastDecimal, _ := decimal.NewFromString(dto.Coast)

		resBalanceDecimal = resBalanceDecimal.Sub(coastDecimal)

		localUser.ResB = resBalanceDecimal.String()

		localOrder := s.repositoryOrder.GetStateByParams(*dto)
		if (order.Order{} != localOrder) {
			if localOrder.State == "reserved" {

				err := s.repositoryUser.UpdateUser(localUser)
				if err != nil {
					return "", err
				}

				err = s.repositoryOrder.UpdateOrder(*dto)
				if err != nil {
					return "", err
				}
				return "Recognition ended successfully", nil
			} else {
				return "", fmt.Errorf("this opetation recognitioned or canceled")
			}
		} else {
			return "", fmt.Errorf("order not found")
		}
	} else {
		return "", fmt.Errorf("user not found")
	}
}
