package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"server04/controller"
	"testing"

	"github.com/gorilla/mux"
)

func TestLikeGetController(t *testing.T) {
	req, err := http.NewRequest("GET", "/like/getVideoLikes/1", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.LikeGetController)
	vars := map[string]string{
		"videoId": "1",
	}

	// CHANGE THIS LINE!!!
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"success":true,"message":"","data":{}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
