package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	render := &Template{
		template: template.Must(template.ParseGlob("public/index.html")),
	}
	e.Renderer = render
	e.POST("/", func(c echo.Context) error {
		name := c.FormValue("message")
		fmt.Println(name)
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
