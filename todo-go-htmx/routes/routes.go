package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"todo-go-htmx/model"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("Could not execute index template", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
	}

	fmt.Println("Todos", todos)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("Could not execute index template", err)
	}
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Marking todo")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating todo")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Could not parse form", err)
	}

	err = model.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Println("Could not create todo", err)
	}

	fmt.Println("printing w from createTodo", w)

	sendTodos(w)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting todo")
}

func SetupAndRun() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/todos/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")

	println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
