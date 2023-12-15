package router

import (
	"net/http"
	"server04/controller"

	"github.com/gorilla/mux"
)

func registerUserRouter(r *mux.Router) {

	userRouter := r.PathPrefix("/user").Subrouter()
	// userRouter.HandleFunc("/", controller.UserGetAllController).Methods(http.MethodGet)
	userRouter.HandleFunc("/pip", controller.UserValidateController).Methods(http.MethodGet)
	userRouter.HandleFunc("/create", controller.UserCreateController).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", controller.UserLoginController).Methods(http.MethodPost)
	// userRouter.HandleFunc("/{id}", controller.UserDeleteController).Methods(http.MethodDelete)
}
