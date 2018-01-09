package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv("DBHost", "mysql")
	os.Setenv("Username", "root")
	os.Setenv("Password", "root")
	os.Setenv("Host", "tcp(172.17.0.2:3306)")
	os.Setenv("DBName", "todoserver")
	os.Setenv("Charset", "utf8")
	os.Setenv("ParseTime", "True")
	os.Setenv("Loc", "Local")

	ListRepo := GormListRepo{}

	fetchList := FetchListIntent{ListRepo}
	createList := CreateListIntent{ListRepo}
	updateList := UpdateListNameIntent{ListRepo}
	deleteList := DeleteListIntent{ListRepo}
	fetchAll := FetchAllListIntent{ListRepo}

	fetchTask := FetchTaskIntent{ListRepo}
	createTask := CreateTaskIntent{ListRepo}
	updateTask := UpdateTaskIntent{ListRepo}
	deleteTask := DeleteTaskIntent{ListRepo}

	router := mux.NewRouter()

	router.HandleFunc("/lists/{id}", fetchList.Enact).Methods("GET")
	router.HandleFunc("/lists", createList.Enact).Methods("POST")
	router.HandleFunc("/lists", updateList.Enact).Methods("PATCH")
	router.HandleFunc("/lists/{id}", deleteList.Enact).Methods("DELETE")
	router.HandleFunc("/lists", fetchAll.Enact).Methods("GET")

	router.HandleFunc("/lists/{id}/tasks/{id}", fetchTask.Enact).Methods("GET")
	router.HandleFunc("/lists/{id}/tasks", createTask.Enact).Methods("POST")
	router.HandleFunc("/lists/{id}/tasks", updateTask.Enact).Methods("PATCH")
	router.HandleFunc("/lists/{id}/tasks/{id}", deleteTask.Enact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
	http.Handle("/", router)
}
