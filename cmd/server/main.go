package main

import (
	"github.com/ayhonz/kaboom/internal/features/home"
	"github.com/ayhonz/kaboom/internal/features/todo"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", home.HomeHandler)
	e.POST("todos", todo.CreateToDoHandler)

	e.Start(":8080")
}
