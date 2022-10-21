package main

import (
	"avitoTechInternship/internal/config"
	"avitoTechInternship/internal/order"
	orderDB "avitoTechInternship/internal/order/db"
	"avitoTechInternship/internal/user"
	userDB "avitoTechInternship/internal/user/db"
	"avitoTechInternship/pkg/client/mysqldb"
	"avitoTechInternship/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	localConfig := config.GetConfig()
	localDB := mysqldb.NewClient(localConfig, logger)

	userRepo := userDB.NewRepository(localDB, logger)
	orderRepo := orderDB.NewRepository(localDB, logger)

	userService := user.NewService(userRepo, logger)
	orderService := order.NewService(orderRepo, logger)

	logger.Info("register handler")

	handler := user.NewHandler(logger, userService, orderService)
	handler.Register(router)

	logger.Info("start server")
	start(router, logger, localConfig)
}

func start(router *httprouter.Router, logger *logging.Logger,
	localConfig *config.Config) {

	var listener net.Listener

	if localConfig.Listen.Type == "port" {
		logger.Info("create tcp server")
		port := localConfig.Listen.Port
		add := localConfig.Listen.BindIp

		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", add, port))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Infof("server listening %s:%s", add, port)
	} else {
		logger.Fatal("error to configure server")
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Fatal(server.Serve(listener))
}
