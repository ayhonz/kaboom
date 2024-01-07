package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ayhonz/kaboom/internal/model"
)


type application struct {
	Todos []model.Todo
}

func main() {
	e := echo.New()

	app := &application{
		Todos: []model.Todo{
			{ID: "1", Text: "todo"},
			{ID: "2", Text: "go outside"},
			{ID: "3", Text: "learn new algos"},
		},
	}

	e.Static("/", "internal/assets")

	e.GET("/", app.HomeHandler)
	e.POST("todos", app.CreateToDoHandler)
	e.DELETE("todos/:id", app.DeleteTodoHandler)

	e.Start(":8080")
}
