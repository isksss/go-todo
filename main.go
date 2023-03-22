package main

import (
	"flag"
	"os"

	"github.com/isksss/go-todo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// _ "github.com/mattn/go-sqlite3"
)

var (
	address string = ":8080"
)

var (
	serverStart bool
)

func flagParser() {
	flag.BoolVar(&serverStart, "server", false, "todo-web start")
}

func newRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// page
	e.Static("/assets", "public/assets")

	e.File("/", "public/index.html")

	// api
	v1 := e.Group("/v1")

	v1.GET("/todo", handler.GetToDos)
	return e
}

func run() int {
	handler.Config()

	flagParser()
	flag.Parse()

	if serverStart {
		r := newRouter()
		r.Logger.Fatal(r.Start(address))
	}

	return 0
}

func main() {
	os.Exit(run())
}
