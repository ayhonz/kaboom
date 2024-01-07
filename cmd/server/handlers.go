package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/ayhonz/kaboom/internal/model"
	"github.com/ayhonz/kaboom/internal/templates/components"
	"github.com/ayhonz/kaboom/internal/templates/pages"
	"github.com/labstack/echo/v4"
)

func (app *application) HomeHandler(c echo.Context) error {
	return pages.HomePage(app.Todos).Render(c.Request().Context(), c.Response())
}

func (app *application) CreateToDoHandler(c echo.Context) error {
	println("creating todo!")
	todo := model.Todo{
		ID:   fmt.Sprintf("%v", len(app.Todos)+1),
		Text: c.FormValue("todo"),
	}
	app.Todos = append(app.Todos, todo)
	fmt.Printf("%+v\n", todo)
	return components.TodoItem(todo).Render(c.Request().Context(), c.Response())
}

func (app *application) DeleteTodoHandler(c echo.Context) error {
	println("deleting todo")

	id := c.Param("id")

	for i, todo := range app.Todos {
		if todo.ID == id {
			app.Todos = slices.Delete(app.Todos, i, i+1)
		}
	}

	return c.NoContent(http.StatusOK)
}
