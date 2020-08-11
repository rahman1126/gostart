package http

import (
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router, h *Handler) {
	r.HandleFunc("/user/{id:[0-9]+}", h.FindOneUser).Methods("GET")
	r.HandleFunc("/users", h.FindAllUsers).Methods("GET")
}
