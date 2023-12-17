package router

import (
	"github.com/gorilla/mux"
)

func RegisterRouter(r *mux.Router) {
	registerVideoRouter(r)
	registerUserRouter(r)
}
