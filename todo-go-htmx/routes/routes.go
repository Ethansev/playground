package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Marking todo")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating todo")
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

	log.Fatal(http.ListenAndServe(":8080", mux))
}
