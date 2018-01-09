package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//FetchTaskIntent is an intent to access task
type FetchTaskIntent struct {
	ListRepo ListRepository
}

//Enact function is for FetchTaskIntent to fetch task through http
func (fetchTask FetchTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		list               List
		dbError, httpError error
	)

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	i, httpError := strconv.Atoi(params["id"])
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}

	list.Tasks = make([]Task, 1)
	list.Tasks[0].ID = uint(i)
	list, dbError = fetchTask.ListRepo.FetchTaskByID(list)

	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusNoContent|http.StatusBadRequest)
	}
	taskJSON, httpError := json.Marshal(list.Tasks[0])
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(taskJSON)
}
