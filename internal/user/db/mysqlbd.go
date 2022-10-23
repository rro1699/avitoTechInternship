package db

import (
	"avitoTechInternship/internal/user"
	"avitoTechInternship/pkg/logging"
	"database/sql"
	"fmt"
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

func (d *databaseOperations) GetUserById(dto user.UserDTO) user.User {
	query, err := d.database.Query("SELECT id, curB, resB FROM users WHERE id = ?", dto.Id)
	if err != nil {
		d.logger.Error("User not found")
		return user.User{}
	}
	if query.Next() {
		var localUser user.User
		err = query.Scan(&localUser.Id, &localUser.CurB, &localUser.ResB)
		if err != nil {
			d.logger.Error("failed mapping data from database to user object")
			return user.User{}
		}
		return localUser
	} else {
		return user.User{}
	}
}

func (d *databaseOperations) UpdateUser(localUser user.User) error {
	_, err := d.database.Query("UPDATE users set curB = ?, resB = ? WHERE id = ?",
		localUser.CurB, localUser.ResB, localUser.Id)
	if err != nil {
		return fmt.Errorf("failed mapping data from user object to database")
	}
	return nil
}

func (d *databaseOperations) CreateUser(localUser user.User) (string, error) {
	result, err1 := d.database.Exec("INSERT INTO users(curB,resB) VALUES (?,?)",
		localUser.CurB, localUser.ResB)
	if err1 != nil {
		return "", fmt.Errorf("failed insert new user")
	}
	id, err1 := result.LastInsertId()
	if err1 != nil {
		return "", fmt.Errorf("failed get id  new user")
	}
	return fmt.Sprint(id), nil
}

/*func (d *databaseOperations) Accrual(dto user.UserDTO) (string, error) {
	d.logger.Debug("Accrual money")
	//getUserById(user UserDTO) User
	query, err := d.database.Query("SELECT id, curB, resB FROM users WHERE id = ?", dto.Id)
	if err != nil {
		return "", fmt.Errorf("failed accurual money")
	}
	if query.Next() {
		var user user.User
		err = query.Scan(&user.Id, &user.CurB, &user.ResB)
		if err != nil {
			return "", fmt.Errorf("failed mapping data from database to user object")
		}
		curBalanceDecimal, _ := decimal.NewFromString(user.CurB)
		coastDecimal, _ := decimal.NewFromString(dto.Coast)
		curBalanceDecimal = curBalanceDecimal.Add(coastDecimal)

		user.CurB = curBalanceDecimal.String()

		//UpdataUser(user User)
		_, err = d.database.Query("UPDATE users set curB = ? WHERE id = ?", user.CurB, user.Id)
		if err != nil {
			return "", fmt.Errorf("failed mapping data from user object to database")
		}
		return fmt.Sprint(user.Id), nil
	} else {
		d.logger.Debug("Insert new user")
		curBalanceDecimal, _ := decimal.NewFromString(dto.Coast)

		//CreateUser(user useDTO)
		result, err1 := d.database.Exec("INSERT INTO users(curB,resB) VALUES (?,?)",
			curBalanceDecimal.String(), decimal.Zero.String())
		if err1 != nil {
			return "", fmt.Errorf("failed insert new user")
		}
		id, err1 := result.LastInsertId()
		if err1 != nil {
			return "", fmt.Errorf("failed get id  new user")
		}
		return fmt.Sprint(id), nil
	}
}
*/

/*func (d *databaseOperations) GetBalance(dto user.UserDTO) (user.User, error) {
	d.logger.Debug("Get balance")
	//getUserById(user UserDTO) User
	query, err := d.database.Query("SELECT id, curB, resB FROM users WHERE id = ?", dto.Id)
	if err != nil {
		return user.User{}, fmt.Errorf("failed get balance")
	}
	var localUser user.User
	if query.Next() {
		err = query.Scan(&localUser.Id, &localUser.CurB, &localUser.ResB)
		if err != nil {
			return user.User{}, fmt.Errorf("failed mapping data from database to object user")
		}
	} else {
		return user.User{}, fmt.Errorf("not found user")
	}
	return localUser, nil
}*/

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
