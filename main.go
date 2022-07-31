package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ToDo struct {
	Title string `json:"Title"`
	Desc  string `json:"desc"`
	Id    int    `json:"id"`
}

// let's declare a global ToDo array
// that we can then populate in our main function
// to simulate a database
var ToDos []ToDo

func returnAllToDos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ToDos)
}

func handleRequests() {
	http.HandleFunc("/todos", returnAllToDos)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func main() {
	ToDos = []ToDo{
		ToDo{Title: "Task 1", Desc: "Make UI", Id: 1},
		ToDo{Title: "Task 2", Desc: "Integrate UI with REST Api", Id: 2},
	}
	handleRequests()
}
