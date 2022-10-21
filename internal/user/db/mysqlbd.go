package db

import (
	"avitoTechInternship/internal/user"
	"avitoTechInternship/pkg/logging"
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
)

type databaseOperations struct {
	database *sql.DB
	logger   *logging.Logger
}

func NewRepository(database *sql.DB, logger *logging.Logger) user.Repository {
	return &databaseOperations{
		database: database,
		logger:   logger,
	}
}

func (d *databaseOperations) Accrual(dto user.UserDTO) error {
	d.logger.Debug("Accrual money")
	query, err := d.database.Query("SELECT id, curB, resB FROM users WHERE id = ?", dto.Id)
	if err != nil {
		return fmt.Errorf("failed accurual money: %v", err)
	}
	if query.Next() {
		var user user.User
		err = query.Scan(&user.Id, &user.CurB, &user.ResB)
		if err != nil {
			return fmt.Errorf("failed mapping data from database to user object: %v", err)
		}
		curBalanceDecimal, _ := decimal.NewFromString(user.CurB)
		coastDecimal, _ := decimal.NewFromString(dto.Coast)
		curBalanceDecimal = curBalanceDecimal.Add(coastDecimal)

		user.CurB = curBalanceDecimal.String()

		_, err = d.database.Query("UPDATE users set curB = ? WHERE id = ?", user.CurB, user.Id)
		if err != nil {
			return fmt.Errorf("failed mapping data from user object to database: %v", err)
		}
		return nil
	} else {
		d.logger.Debug("Insert new user")
		curBalanceDecimal, _ := decimal.NewFromString(dto.Coast)
		_, err = d.database.Exec("INSERT INTO users(curB,resB) VALUES (?,?)",
			curBalanceDecimal.String(), decimal.Zero.String())
		if err != nil {
			return fmt.Errorf("failed insert new user: %v", err)
		}
		return nil
	}
}

func (d *databaseOperations) GetBalance(dto user.UserDTO) (user.User, error) {
	d.logger.Debug("Get balance")
	query, err := d.database.Query("SELECT id, curB, resB FROM users WHERE id = ?", dto.Id)
	if err != nil {
		return user.User{}, fmt.Errorf("failed get balance: %v", err)
	}
	var localUser user.User
	if query.Next() {
		err = query.Scan(&localUser.Id, &localUser.CurB, &localUser.ResB)
		if err != nil {
			return user.User{}, fmt.Errorf("failed mapping data from database to object user: %v", err)
		}
	} else {
		return user.User{}, fmt.Errorf("not found user: %v", err)
	}
	return localUser, nil
}

/*

	curBalanceDecimal, _ := decimal.NewFromString(user.CurB)
	resBalanceDecimal, _ := decimal.NewFromString(user.ResB)
	coastDecimal, _ := decimal.NewFromString(dto.Coast)

	if sub := curBalanceDecimal.Sub(coastDecimal); sub.Cmp(decimal.Zero) < 0 {
		return fmt.Errorf("you don't have enough money on balance")
	} else {
		resBalanceDecimal = resBalanceDecimal.Add(coastDecimal)
		curBalanceDecimal = curBalanceDecimal.Sub(coastDecimal)
		user.CurB = curBalanceDecimal.String()
		user.ResB = resBalanceDecimal.String()
		_, err = d.database.Query("UPDATE users set curB = ?, resB = ? WHERE id = ?", user.CurB, user.ResB, user.Id)
		if err != nil {
			return fmt.Errorf("failed mapping data from user object to database: %s", err)
		}
		return nil
	}
*/
