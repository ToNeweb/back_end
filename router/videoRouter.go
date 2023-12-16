package router

import (
	"net/http"
	"server04/controller"

	"github.com/gorilla/mux"
)

func registerVideoRouter(r *mux.Router) {

	videoRouter := r.PathPrefix("/video").Subrouter()
	videoRouter.HandleFunc("/putvideo", controller.VideoPutController).Methods(http.MethodPost)
	// todoRouter.HandleFunc("/", controller.TodoCreateController).Methods(http.MethodPost)
	videoRouter.HandleFunc("/getvideo/{id}", controller.VideoGetController).Methods(http.MethodGet)
	videoRouter.HandleFunc("/getvideos", controller.VideoGetBatchController).Methods(http.MethodGet)
}

///eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTEyLTE1VDE4OjAwOjI1LjQ5MjMwNjIrMDM6MzAiLCJmb28iOiJiYXIiLCJ1c2VySWQiOjF9.WbSyMGClrOzOn1SZ8qLqtWM5J-zqk_hWnAK9vypEjJU
