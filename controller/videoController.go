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

// func VideoPutController(w http.ResponseWriter, r *http.Request) { /// this feature is hard to test
// 	var headersOfFiles [2]string
// 	if r.Method != "POST" {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// 32 MB is the default used by FormFile()
// 	if err := r.ParseMultipartForm(32 << 20); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Get a reference to the fileHeaders.
// 	// They are accessible only after ParseMultipartForm is called
// 	files := r.MultipartForm.File["file"]

// 	for i, fileHeader := range files {
// 		if i == 2 {
// 			break
// 		}

// 		// Restrict the size of each uploaded file to 1MB.
// 		// To prevent the aggregate size from exceeding
// 		// a specified value, use the http.MaxBytesReader() method
// 		// before calling ParseMultipartForm()
// 		var MAX_UPLOAD_SIZE int64 = 10000000000 /// CHANGE THIS LATER
// 		if fileHeader.Size > MAX_UPLOAD_SIZE {
// 			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
// 			return
// 		}

// 		// Open the file
// 		file, err := fileHeader.Open()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		defer file.Close()

// 		buff := make([]byte, 512)
// 		_, err = file.Read(buff)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		_, err = file.Seek(0, io.SeekStart)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		err = os.MkdirAll("./uploads", os.ModePerm)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		headersOfFiles[i] = fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
// 		f, err := os.Create(headersOfFiles[i])
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		defer f.Close()

// 		_, err = io.Copy(f, file)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 	}

// 	fmt.Fprintf(w, "Upload successful")

// 	var newVideo ent.Videos
// 	json.NewDecoder(r.Body).Decode(&newVideo)
// 	newVideo.Thumb = headersOfFiles[0]
// 	newVideo.VideoLink = headersOfFiles[1]

// 	video, err := service.NewvideosOps(r.Context()).VideoCreate(newVideo, (int)(r.Context().Value("user_id").(float64)))
// 	fmt.Println(err)
// 	fmt.Println("video is ", video)
// 	utils.Return(w, true, http.StatusOK, nil, video)

// }

func VideoPutController(w http.ResponseWriter, r *http.Request) { /// https://freshman.tech/file-upload-golang/

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

func VideoGetBatchController(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	minLast, _ := strconv.Atoi(v["min_last"][0])
	numberRequested, _ := strconv.Atoi(v["number_requested"][0])
	videos, _ := service.NewvideosOps(r.Context()).VideoGetBatch(minLast, numberRequested)
	utils.Return(w, true, http.StatusOK, nil, videos)
}
func VideoGetSearchController(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	titleSearch := v["title"][0]

	videos, _ := service.NewvideosOps(r.Context()).VideoGetSearch(titleSearch)

	utils.Return(w, true, http.StatusOK, nil, videos)
}
