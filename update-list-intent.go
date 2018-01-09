package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//UpdateListNameIntent is an intent to update single list name
type UpdateListNameIntent struct {
	ListRepo ListRepository
}

//Enact function is for UpdateListNameIntent to update list through http
func (updateListIntent UpdateListNameIntent) Enact(w http.ResponseWriter, r *http.Request) {
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

	dbError = updateListIntent.ListRepo.Update(list)
	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusNoContent|http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
