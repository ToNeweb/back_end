package router

import (
	"net/http"
	"server04/controller"

	"github.com/gorilla/mux"
)

func registerUserRouter(r *mux.Router) {

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/pip", controller.UserExpireController).Methods(http.MethodGet)
	userRouter.HandleFunc("/create", controller.UserCreateController).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", controller.UserLoginController).Methods(http.MethodPost)
	userRouter.HandleFunc("/validate/createwithvalidate", controller.UserValidationSendController).Methods(http.MethodPost)
	userRouter.HandleFunc("/validate/validation", controller.UserValidationCheckController).Methods(http.MethodPost)
}
