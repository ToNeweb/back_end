package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server04/controller"
	"server04/ent"
	"testing"
)

func TestVideoPutController(t *testing.T) {
	// Create a new video object
	newVideo := &ent.Videos{
		Desc: "Test Video",
		// Add other fields as needed
	}

	// Marshal the video object to JSON
	videoJSON, err := json.Marshal(newVideo)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the method "PUT" and the URL "/video"
	req, err := http.NewRequest("PUT", "/video", bytes.NewBuffer(videoJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Set the content type of the request to "application/json"
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.VideoPutController)

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
