package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(r *mux.Router, h *Handler) {
	r.HandleFunc("/api/health", h.HealthCheck)
	r.Use(mux.CORSMethodMiddleware(r))

	auth := r.PathPrefix("/api/v1").Subrouter()
	auth.HandleFunc("/user/{id:[0-9]+}", h.FindOneUser).Methods(http.MethodGet, http.MethodOptions)
	auth.HandleFunc("/users", h.FindAllUsers).Methods(http.MethodGet, http.MethodOptions)
	auth.Use(AuthMiddleware)
	auth.Use(mux.CORSMethodMiddleware(auth))
}