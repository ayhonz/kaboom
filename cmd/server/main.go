package main

import (
	"github.com/ayhonz/kaboom/internal/features/home"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", home.HomeHandler)

	e.Start(":8080")
}
