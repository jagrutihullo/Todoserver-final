package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//UpdateTaskIntent is an intent to update single task
type UpdateTaskIntent struct {
	ListRepo ListRepository
}

//Enact function is for UpdateTaskIntent to update task through http
func (updateTask UpdateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		task               Task
		dbError, httpError error
		list               List
	)

	w.Header().Set("Content-Type", "application/json")
	body, httpError := ioutil.ReadAll(r.Body)
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}

	httpError = json.Unmarshal(body, &task)
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}

	list.Tasks = make([]Task, 1)
	list.Tasks[0] = task
	dbError = updateTask.ListRepo.UpdateTask(list)
	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusNoContent|http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
