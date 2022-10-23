package user

import (
	"avitoTechInternship/internal/order"
	"avitoTechInternship/internal/report"
	"avitoTechInternship/pkg/logging"
	"encoding/csv"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
)

type Service struct {
	repositoryOrder  order.Repository
	repositoryReport report.Repository
	repositoryUser   Repository
	logger           *logging.Logger
}

func NewService(repositoryOrder order.Repository, repositoryReport report.Repository,
	repositoryUser Repository, logger *logging.Logger) *Service {
	return &Service{
		repositoryOrder:  repositoryOrder,
		repositoryReport: repositoryReport,
		repositoryUser:   repositoryUser,
		logger:           logger,
	}
}

func (s *Service) Accrual(dto *UserDTO) (string, error) {
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
	localUser := s.repositoryUser.GetUserById(*dto)
	if (User{} != localUser) {
		return localUser, nil
	} else {
		return User{}, fmt.Errorf("user not found")
	}
}

func (s *Service) Reservation(dto *order.OrderDTO) (string, error) {

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

func (s *Service) GetReport(dto *report.ReportDTO) (string, error) {
	if dto.Year <= 0 {
		return "", fmt.Errorf("badly number of year")
	}
	lastDay := findLastDayOfMonth(dto.Year, dto.Month)
	if lastDay == -1 {
		return "", fmt.Errorf("badly number of month")
	}
	orders := s.repositoryOrder.GetOrdersByDate(dto.Year, dto.Month, lastDay)
	if orders != nil && len(orders) > 0 {
		reportInfo := make(map[int]decimal.Decimal)
		for _, val := range orders {
			coastDecimal, _ := decimal.NewFromString(val.Coast)
			curCoastDecimal, _ := reportInfo[val.IdSer]

			curCoastDecimal = curCoastDecimal.Add(coastDecimal)

			reportInfo[val.IdSer] = curCoastDecimal
		}
		dataCSV := make(map[string]string)
		for key, val := range reportInfo {
			servInfo := s.repositoryReport.GetServiceById(key)
			if (report.Serv{} != servInfo) {
				dataCSV[servInfo.Name] = val.String()
			} else {
				return "", fmt.Errorf("service not found")
			}
		}
		filePath, err := generationCsv(dataCSV, dto.Year, dto.Month)
		if err != nil {
			return "", err
		}
		return filePath, nil
	} else {
		return "", fmt.Errorf("don't have orders for this period")
	}
}

func generationCsv(data map[string]string, year, month int) (string, error) {
	pathDir := fmt.Sprintf("./%s", "csvdata")
	_, err := os.Stat(pathDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(pathDir, 0640)
		if err != nil {
			return "", fmt.Errorf("problems with create dir")
		}
	}
	filePath := fmt.Sprintf("%s/report%d_%d.csv", pathDir, month, year)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	defer file.Close()

	if err != nil {
		return "", fmt.Errorf("problems with opening file %s", filePath)
	}

	w := csv.NewWriter(file)
	err = w.Write([]string{"название услуги", "общая сумма выручки за отчетный период"})
	if err != nil {
		return "", fmt.Errorf("problems with writing into file %s", filePath)
	}
	for name, coast := range data {
		if err = w.Write([]string{name, coast}); err != nil {
			return "", fmt.Errorf("problems with writing into file %s", filePath)
		}
	}
	w.Flush()
	if err = w.Error(); err != nil {
		return "", fmt.Errorf("problems with closing file %s", filePath)
	}

	return file.Name(), nil
}

func findLastDayOfMonth(year, month int) int {
	flag := year%4 == 0 && year%100 != 0 || year%400 == 0
	switch month {
	case 1:
		return 31
	case 2:
		if flag {
			return 29
		} else {
			return 28
		}
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	case 12:
		return 31
	default:
		return -1
	}
}
