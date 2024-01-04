package handlers

import (
	"github.com/ayhonz/kaboom/internal/templates/pages"
	"github.com/labstack/echo/v4"
)

var todos = []string{"one", "two"}

func HomeHandler(c echo.Context) error {
	return pages.HomePage(todos).Render(c.Request().Context(), c.Response())
}
