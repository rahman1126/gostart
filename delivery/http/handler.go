package http

import (
	"context"
	"github.com/gorilla/mux"
	"gostart/usecase"
	"gostart/utils/common"
	"gostart/utils/logger"
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

func (h Handler) FindAllUsers(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	data, err := h.ExampleUsecase.FindAll(ctx)

	common.Response(w, data, err)
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

	common.Response(w, data, err)
}