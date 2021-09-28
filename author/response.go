package author

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SetResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	AccessTime string      `json:"accessTime"`
}

func ResponseSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	setResponse := SetResponse{
		Status:     http.StatusText(200),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data}
	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, r *http.Request, data interface{}, code int, status error) {
	setResponse := SetResponse{
		Status:     fmt.Sprintf("%v", status),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data}
	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}
