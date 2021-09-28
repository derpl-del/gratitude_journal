package author

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var a Author
	json.Unmarshal(ReqBody, &a)
	err := a.Create()
	if err != nil {
		ResponseError(w, r, a, 500, err)
		return
	}
	ResponseSuccess(w, r, a)
}

func QueryAuthor(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var a Author
	json.Unmarshal(ReqBody, &a)
	err := a.Get()
	if err != nil {
		ResponseError(w, r, a, 500, err)
		return
	}
	ResponseSuccess(w, r, a)
}

func UpdateJournal(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var a Author
	json.Unmarshal(ReqBody, &a)
	err := a.Update()
	if err != nil {
		ResponseError(w, r, a, 500, err)
		return
	}
	ResponseSuccess(w, r, a)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var a Author
	json.Unmarshal(ReqBody, &a)
	err := a.Delete()
	if err != nil {
		ResponseError(w, r, a, 500, err)
		return
	}
	ResponseSuccess(w, r, a)
}
