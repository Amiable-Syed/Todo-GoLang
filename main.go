package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func createNewToDo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	// get the body of our POST request
	// unmarshal this into a new ToDo struct
	// append this to our ToDos array.
	var todo ToDo
	json.Unmarshal(reqBody, &todo)
	//update our global ToDo Array to include our new ToDo
	ToDos = append(ToDos, todo)

	json.NewEncoder(w).Encode(todo)
}

func deleteToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	//Loop on all the ToDos
	// get the matched Id and remove
	// the object from the global array
	for idx, todo := range ToDos {
		if todo.Id == key {
			ToDos = append(ToDos[:idx], ToDos[idx+1:]...)
		}
	}

}

func handleRequests() {
	// Mux Router introduced instead traditional net/http router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/todos", returnAllToDos)
	myRouter.HandleFunc("/todo/{id}", returnSingleToDo)
	myRouter.HandleFunc("/todo/{id", deleteToDo).Methods("DELETE")
	myRouter.HandleFunc("todo", createNewToDo).Methods("POST")
	log.Fatal(http.ListenAndServe(":1000", myRouter))
}

func main() {
	ToDos = []ToDo{
		ToDo{Title: "Task 1", Desc: "Make UI", Id: "1"},
		ToDo{Title: "Task 2", Desc: "Integrate UI with REST Api", Id: "2"},
	}
	handleRequests()
}
