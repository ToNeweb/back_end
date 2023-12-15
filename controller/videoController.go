package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server04/ent"
	"server04/service"
	"server04/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func VideoPutController(w http.ResponseWriter, r *http.Request) {

	var newVideo ent.Videos
	json.NewDecoder(r.Body).Decode(&newVideo)
	// if err != nil {
	// 	utils.Return(w, false, http.StatusBadRequest, err, nil)
	// 	return
	// }
	video, err := service.NewvideosOps(r.Context()).VideoCreate(newVideo, (int)(r.Context().Value("user_id").(float64)))
	fmt.Println(err)
	//service.NewVideosClient(r.Context()).Video
	//user, err := service.newVideo(r.Context()).UserCreate(newUser)
	fmt.Println("video is ", video)
	utils.Return(w, true, http.StatusOK, nil, video)
	//fmt.Println("user ", r.Context().Value("user_id"), " wants to put video")
}

func VideoGetController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	todo, err := service.NewvideosOps(r.Context()).VideoGetByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, todo)
}
