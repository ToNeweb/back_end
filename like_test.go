package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestLikeGetController(t *testing.T) {
	// // Create a request to use for our handler. We don't have any query parameters for now, so we'll
	// // pass 'nil' as the third parameter.
	// req, err := http.NewRequest("GET", "/like/getVideoLikes/1", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(controller.LikeGetController)
	// vars := map[string]string{
	// 	"id": "1",
	// }

	// // CHANGE THIS LINE!!!
	// req = mux.SetURLVars(req, vars)
	// // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// // directly and pass in our Request and ResponseRecorder.
	// handler.ServeHTTP(rr, req)
	rr, _ := http.Get("http://localhost:8019/like/getVideoLikes/1")

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Check the status code is what we expect.
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

	if string(body)[44] != '1' {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), 1)
	}
}
