package controller

import (
	"log"
	"net/http"
	"server04/service"
	"server04/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func LikePutController(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	videoId, _ := strconv.Atoi(v["videoId"][0])
	isNotOk := service.NewvideosOps(r.Context()).				AddVideoLikeIfVideoDoesNotHaveLikeFromUser(videoId, (int)(r.Context().Value("user_id").(float64)))
	if !isNotOk {
		like, err := service.NewLikesOps(r.Context()).PutLike((int)(r.Context().Value("user_id").(float64)), videoId)
		log.Println(err)
		utils.Return(w, true, http.StatusOK, nil, like)
	} else {
		utils.Return(w, true, http.StatusOK, nil, false)
	}
}
func LikeGetController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId, _ := strconv.Atoi(vars["id"])
	videos, _ := service.NewvideosOps(r.Context()).VideoGetByID(videoId)

	utils.Return(w, true, http.StatusOK, nil, videos.LikeNum)
}
