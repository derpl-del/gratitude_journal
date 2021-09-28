package journal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateJournal(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var j Journal
	json.Unmarshal(ReqBody, &j)
	err := j.Create()
	if err != nil {
		ResponseError(w, r, j, 500, err)
		return
	}
	ResponseSuccess(w, r, j)
}

func QueryJournal(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var j Journal
	json.Unmarshal(ReqBody, &j)
	err := j.Get()
	if err != nil {
		ResponseError(w, r, j, 500, err)
		return
	}
	ResponseSuccess(w, r, j)
}

func UpdateJournal(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var j Journal
	json.Unmarshal(ReqBody, &j)
	err := j.Update()
	if err != nil {
		ResponseError(w, r, j, 500, err)
		return
	}
	ResponseSuccess(w, r, j)
}

func DeleteJournal(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var j Journal
	json.Unmarshal(ReqBody, &j)
	err := j.Delete()
	if err != nil {
		ResponseError(w, r, j, 500, err)
		return
	}
	ResponseSuccess(w, r, j)
}
