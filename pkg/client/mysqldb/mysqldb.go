package mysqldb

import (
	"avitoTechInternship/internal/config"
	"avitoTechInternship/pkg/logging"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewClient(localConfig *config.Config, logger *logging.Logger) *sql.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", localConfig.Database.User, localConfig.Database.Password,
		localConfig.Database.Host, localConfig.Database.Port, localConfig.Database.DbName)
	open, err := sql.Open("mysql", dbURL)

	if err != nil {
		logger.Fatal("Connection to database ended not success")
		return nil
	}
	if err = open.Ping(); err != nil {
		logger.Fatal("Ping database ended badly")
		return nil
	}
	logger.Info("Successfully connection to database")
	return open
}
