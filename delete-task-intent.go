package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//DeleteTaskIntent is an intent to delete task
type DeleteTaskIntent struct {
	ListRepo ListRepository
}

//Enact function is for DeleteTaskIntent to delete task through http
func (deleteTask DeleteTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
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

	list.Tasks = make([]Task, 1)
	list.Tasks[0].ID = uint(i)
	dbError = deleteTask.ListRepo.DeleteTask(list)
	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
