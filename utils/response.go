package utils

import (
	"encoding/json"
	"net/http"
	"server04/model"
)

func Return(w http.ResponseWriter, status bool, code int, err error, data interface{}) {

	response := model.Response{
		Status: status,
		Code:   code,
		Error:  "",
		Data:   data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}
