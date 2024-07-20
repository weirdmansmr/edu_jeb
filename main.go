package main

import (
	"education/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    InitDB()
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/lecture/:id", handlers.GetLecture(DB))
    e.Start(":8080")
}

