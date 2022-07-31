# Todo-GoLang
Rest Apis for Todo Application in Go Lang

Apis include all the CRUD operations except put.

## Routes:
/                 -> Home Page

/todos            -> Get All the ToDos  (GET) 

/todo/{id}        -> Get Single ToDo based on ID  (GET) 

/todo/{id}        -> Delete Single ToDo based on ID  (DELETE) 

/todo             -> Create new ToDo (POST)

### Sample Json for Post Call:
{
  Title: "Title", 
  Desc: "Description", 
  Id: "2"
 }

### Dependecies:
Mux/Router
