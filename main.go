package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ToDo struct {
	Title string `json:"Title"`
	Desc  string `json:"desc"`
	Id    string `json:"id"`
}

// let's declare a global ToDo array
// that we can then populate in our main function
// to simulate a database
var ToDos []ToDo

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the Home Page!")
}
func returnAllToDos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ToDos)
}

func returnSingleToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	//Loop on all the ToDos
	// get the matched Id and return
	// the object encoded as json
	for _, todo := range ToDos {
		if todo.Id == key {
			json.NewEncoder(w).Encode(todo)
		}
	}

}
func handleRequests() {
	// Mux Router introduced instead traditional net/http router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/todos", returnAllToDos)
	myRouter.HandleFunc("/todo/{id}", returnSingleToDo)
	log.Fatal(http.ListenAndServe(":1000", myRouter))
}

func main() {
	ToDos = []ToDo{
		ToDo{Title: "Task 1", Desc: "Make UI", Id: "1"},
		ToDo{Title: "Task 2", Desc: "Integrate UI with REST Api", Id: "2"},
	}
	handleRequests()
}
