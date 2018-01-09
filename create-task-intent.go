package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateTaskIntent is an intent to create task under list
type CreateTaskIntent struct {
	ListRepo ListRepository
}

//Enact function is for CreateTaskIntent to create task through http
func (createTask CreateTaskIntent) Enact(w http.ResponseWriter, r *http.Request) {
	var (
		list               List
		task               Task
		dbError, httpError error
	)

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, httpError := strconv.Atoi(params["lid"])
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}
	list.ID = uint(i)

	body, httpError := ioutil.ReadAll(r.Body)
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}

	httpError = json.Unmarshal(body, &task)
	if httpError != nil {
		http.Error(w, httpError.Error(), http.StatusBadRequest)
	}
	if task.Description == "" {
		http.Error(w, "Task description cannot be empty", http.StatusBadRequest)
		return
	}

	list.Tasks = make([]Task, 1)
	list.Tasks[0] = task
	dbError = createTask.ListRepo.CreateTask(list)
	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
