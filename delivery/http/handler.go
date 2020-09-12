package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gostart/dto"
	"gostart/usecase"
	"gostart/utils/cache"
	"gostart/utils/common"
	"gostart/utils/conf"
	"gostart/utils/db"
	"gostart/utils/logger"
	"io"
	"net/http"
	"strconv"
)

type Handler struct {
	ExampleUsecase usecase.ExampleUsecaseI
}

func NewHandler(uu usecase.ExampleUsecaseI, r *mux.Router) {
	h := &Handler{
		ExampleUsecase: uu,
	}
	Routes(r, h)
}

func (h Handler) HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	health := &dto.HealthCheckResponse{}
	health.Alive = "true"
	health.Database = "ok"
	health.Redis = "ok"

	err := db.DBConn().DB().Ping()
	if err != nil {
		health.Database = fmt.Sprintf("%v", err)
	}

	redisConn := cache.Pool()
	if conf.IsUsingRedis() {
		err = cache.Ping(redisConn)
		if err != nil {
			health.Redis = fmt.Sprintf("%v", err)
		}
	} else {
		health.Redis = "off"
	}

	statuses, _ := json.Marshal(&health)

	io.WriteString(w, string(statuses))
}

func (h Handler) FindAllUsers(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	data, err := h.ExampleUsecase.FindAll(ctx)

	common.Response(w, req, data, err)
}

func (h Handler) FindOneUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		logger.Error(nil, "FindOneUser: Error converting param ID to integer")
	}

	data, err := h.ExampleUsecase.Find(ctx, id)

	common.Response(w, req, data, err)
}