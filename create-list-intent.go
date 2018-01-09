package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CreateListIntent is an intent to create list
type CreateListIntent struct {
	ListRepo ListRepository
}

//Enact function is for CreateListIntent to create list through http
func (createListIntent CreateListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		list               List
		dbError, httpError error
	)

	w.Header().Set("Content-Type", "application/json")
	body, httpError := ioutil.ReadAll(r.Body)
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}

	httpError = json.Unmarshal(body, &list)
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}

	dbError = createListIntent.ListRepo.Create(list)
	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
