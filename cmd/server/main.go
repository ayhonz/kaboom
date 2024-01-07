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
			model.Todo{ID: "1", Text: "todo"},
			model.Todo{ID: "2", Text: "go outside"},
			model.Todo{ID: "3", Text: "learn new algos"},
		},
	}

	e.GET("/", app.HomeHandler)
	e.POST("todos", app.CreateToDoHandler)
	e.DELETE("todos/:id", app.DeleteTodoHandler)

	e.Start(":8080")
}
