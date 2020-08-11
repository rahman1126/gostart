package main

import (
	"github.com/gorilla/mux"
	delivery "gostart/delivery/http"
	"gostart/repository"
	"gostart/server"
	"gostart/usecase"
	"gostart/utils/cache"
	"gostart/utils/conf"
	"gostart/utils/db"
	"gostart/utils/logger"
)

func init() {
	conf.SetConfigFile("config", "config", "json")
	logger.SetJSONLogger()
}

func main() {
	dbConn := db.DBConn()
	defer dbConn.Close()

	rConn := cache.Conn()
	if rConn != nil {
		defer rConn.Close()
	}

	r := mux.NewRouter()

	er := repository.NewExampleRepository(dbConn)
	eu := usecase.NewExampleUsecase(er, conf.GetCtxTimeout(), rConn)
	delivery.NewHandler(eu, r)

	server.GracefullyShutdown(r)
}