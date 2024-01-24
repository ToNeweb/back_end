package router

import (
	"net/http"
	"server04/controller"

	"github.com/gorilla/mux"
)

func registerLikeRouter(r *mux.Router) {

	videoRouter := r.PathPrefix("/like").Subrouter()
	videoRouter.HandleFunc("/likeVideo", controller.LikePutController).Methods(http.MethodGet)
	videoRouter.HandleFunc("/getVideoLikes/{id}", controller.LikeGetController).Methods(http.MethodGet)
	/// add get likers
}
