package main

import (
	"todo-go-htmx/model"
	"todo-go-htmx/routes"
)

func main() {
	model.Setup()
	routes.SetupAndRun()
}
