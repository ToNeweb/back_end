package router

import (
	"net/http"
	"server04/controller"

	"github.com/gorilla/mux"
)

func registerVideoRouter(r *mux.Router) {

	videoRouter := r.PathPrefix("/video").Subrouter()
	videoRouter.HandleFunc("/putvideo", controller.VideoPutController).Methods(http.MethodPost)
	videoRouter.HandleFunc("/getvideo/{id}", controller.VideoGetController).Methods(http.MethodGet)
	videoRouter.HandleFunc("/getvideos", controller.VideoGetBatchController).Methods(http.MethodGet)
	videoRouter.HandleFunc("/getvideosSearch", controller.VideoGetBatchController).Methods(http.MethodGet)
}
