package main

import (
	"encoding/json"
	"net/http"
)

//FetchAllListIntent is an intent to access all lists
type FetchAllListIntent struct {
	ListRepo ListRepository
}

//Enact function is for FetchAllListIntent to access lists through http
func (fetchAllIntent FetchAllListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		lists              []List
		dbError, httpError error
	)

	lists, dbError = fetchAllIntent.ListRepo.FetchAll()

	w.Header().Set("Content-Type", "application/json")
	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusNoContent|http.StatusBadRequest)
	}
	listsJSON, httpError := json.Marshal(lists)

	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(listsJSON)
}
