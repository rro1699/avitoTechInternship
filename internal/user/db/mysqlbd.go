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
