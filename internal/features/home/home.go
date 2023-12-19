package home

import (
	"github.com/ayhonz/kaboom/internal/templates/pages"
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return pages.HomePage().Render(c.Request().Context(), c.Response())
}
