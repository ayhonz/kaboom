package todo

import (
	"github.com/ayhonz/kaboom/internal/templates/components"
	"github.com/labstack/echo/v4"
)

func CreateToDoHandler(c echo.Context) error {
	println("creating todo!")
	todo := c.FormValue("todo")
	println(todo)
	return components.TodoItem(todo).Render(c.Request().Context(), c.Response())
}
