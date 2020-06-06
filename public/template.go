package public

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func NewRenderer() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("public/**/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
