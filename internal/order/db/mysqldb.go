package db

import (
	"avitoTechInternship/internal/order"
	"avitoTechInternship/pkg/logging"
	"database/sql"
	"fmt"
	"time"
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
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%0.2d", now.Year(), now.Month(), now.Day())
	_, err := d.database.Exec("INSERT INTO orders(idUser,idSer,idOrder,coast, state,orderDate) VALUES (?,?,?,?,?,?)",
		dto.IdUser, dto.IdSer, dto.IdOrder, dto.Coast, "reserved", date)
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

func (d *databaseOperations) GetOrdersByDate(year, month, lastDay int) []order.Order {
	//IdSer,Coast,Date
	startDate := fmt.Sprintf("%d-%d-%.2d", year, month, 1)
	endDate := fmt.Sprintf("%d-%d-%.2d", year, month, lastDay)

	query, err := d.database.Query("SELECT id,idSer,coast FROM orders WHERE state = ? AND (orderDate BETWEEN ? AND ?)",
		"recognition", startDate, endDate)
	if err != nil {
		return nil
	}
	orders := make([]order.Order, 0)
	var localOrder order.Order
	for query.Next() {
		err = query.Scan(&localOrder.Id, &localOrder.IdSer, &localOrder.Coast)
		if err != nil {
			return nil
		}
		orders = append(orders, localOrder)
	}
	return orders
}
