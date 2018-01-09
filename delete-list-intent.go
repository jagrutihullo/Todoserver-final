package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//DeleteListIntent is an intent to delete list & tasks under it
type DeleteListIntent struct {
	ListRepo ListRepository
}

//Enact function is for DeleteListIntent to delete list through http
func (deleteListIntent DeleteListIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		dbError, httpError error
		list               List
	)

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, httpError := strconv.Atoi(params["id"])
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}

	list.ID = uint(i)
	dbError = deleteListIntent.ListRepo.Delete(list)
	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
