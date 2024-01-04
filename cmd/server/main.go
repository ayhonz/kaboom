package main

import (
	"github.com/ayhonz/kaboom/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.HomeHandler)
	e.POST("todos", handlers.CreateToDoHandler)

	e.Start(":8080")
}
