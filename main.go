package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/index.html")
	return e
}

func run() int {
	router := newRouter()

	router.Logger.Fatal(router.Start(":8080"))
	return 0
}

func main() {
	os.Exit(run())
}
