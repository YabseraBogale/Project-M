package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "public/index.html")

	e.Logger.Fatal(e.Start("127.0.0.1:1234"))
}
