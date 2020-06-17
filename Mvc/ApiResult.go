package Mvc

import (
	"encoding/json"
	"io"
	"net/http"
)

type ApiResult struct {
	Success bool
	Message string
	Data    interface{}
}

func Ok(w http.ResponseWriter, data string) {
	io.WriteString(w, data)
}

func Fail(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

func JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(data)
	w.Write(json)
}
