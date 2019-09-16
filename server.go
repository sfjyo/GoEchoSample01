package main

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	render := &TemplateRender {
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.Renderer = render
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
	// e.GET("/users/:name", func(c echo.Context) error {
	// 	name := c.Param("name")
	// 	return c.String(http.StatusOK, "Hello: " + name)
	// })

	e.Logger.Fatal(e.Start(":1323"))
}

