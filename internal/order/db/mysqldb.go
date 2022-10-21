package db

import (
	"avitoTechInternship/internal/order"
	"avitoTechInternship/pkg/logging"
	"database/sql"
	"fmt"
)

type databaseOperations struct {
	database *sql.DB
	logger   *logging.Logger
}

func NewRepository(database *sql.DB, logger *logging.Logger) order.Repository {
	return &databaseOperations{
		database: database,
		logger:   logger,
	}
}

func (d *databaseOperations) Reservation(dto order.OrderDTO) error {
	d.logger.Debug("Reservation money")
	_, err := d.database.Exec("INSERT INTO orders(idUser,idSer,idOrder,coast, state) VALUES (?,?,?,?,?)",
		dto.IdUser, dto.IdSer, dto.IdOrder, dto.Coast, "reserved")
	if err != nil {
		return fmt.Errorf("failed reservation money: %v", err)
	}
	return nil
}

func (d *databaseOperations) Recognition(dto order.OrderDTO) error {
	d.logger.Debug("Recognition order")
	_, err := d.database.Query("UPDATE orders set state = ? WHERE idUser = ?,idSer = ?,idOrder = ?,coast = ?",
		"recognition", dto.IdUser, dto.IdSer, dto.IdOrder, dto.Coast)
	if err != nil {
		return fmt.Errorf("failed recognition order: %v", err)
	}
	return nil
}
