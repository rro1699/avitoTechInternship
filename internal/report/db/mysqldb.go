package db

import (
	"avitoTechInternship/internal/report"
	"avitoTechInternship/pkg/logging"
	"database/sql"
)

type databaseOperations struct {
	database *sql.DB
	logger   *logging.Logger
}

func NewRepository(database *sql.DB, logger *logging.Logger) report.Repository {
	return &databaseOperations{
		database: database,
		logger:   logger,
	}
}

func (d databaseOperations) GetServiceById(idSer int) report.Serv {
	query, err := d.database.Query("SELECT id, name FROM servs WHERE id = ?", idSer)
	if err != nil {
		d.logger.Error("Service not found")
		return report.Serv{}
	}
	if query.Next() {
		var localServ report.Serv
		err = query.Scan(&localServ.IdSer, &localServ.Name)
		if err != nil {
			d.logger.Error("failed mapping data from database to service object")
			return report.Serv{}
		}
		return localServ
	} else {
		return report.Serv{}
	}
}
