package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//FetchListIntent is an intent to access list by name
type FetchListIntent struct {
	ListRepo ListRepository
}

//Enact function is for FetchListIntent to access list through http
func (fetchListIntent FetchListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		list               List
		dbError, httpError error
	)

	params := mux.Vars(r)
	i, httpError := strconv.Atoi(params["id"])
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}

	list.ID = uint(i)
	list, dbError = fetchListIntent.ListRepo.FetchByID(list)

	w.Header().Set("Content-Type", "application/json")
	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusNoContent|http.StatusBadRequest)
	}
	listJSON, httpError := json.Marshal(list)

	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(listJSON)
}
