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

func (d *databaseOperations) Create(dto order.OrderDTO) error {
	_, err := d.database.Exec("INSERT INTO orders(idUser,idSer,idOrder,coast, state) VALUES (?,?,?,?,?)",
		dto.IdUser, dto.IdSer, dto.IdOrder, dto.Coast, "reserved")
	if err != nil {
		return fmt.Errorf("failed reservation money")
	}
	return nil
}

func (d *databaseOperations) GetStateByParams(dto order.OrderDTO) order.Order {
	row := d.database.QueryRow("SELECT state FROM orders WHERE idUser = ? AND idSer = ? AND idOrder = ? AND coast = ?",
		dto.IdUser, dto.IdSer, dto.IdOrder, dto.Coast)
	var localOrder order.Order
	err := row.Scan(&localOrder.State)
	if err != nil {
		return order.Order{}
	}
	return localOrder
}

func (d *databaseOperations) UpdateOrder(dto order.OrderDTO) error {
	_, err := d.database.Query("UPDATE orders SET state = ? WHERE idUser = ? AND idSer = ? AND idOrder = ? AND coast = ?",
		"recognition", dto.IdUser, dto.IdSer, dto.IdOrder, dto.Coast)
	if err != nil {
		return fmt.Errorf("failed recognition order")
	}
	return nil
}
