package constant

// a package to get html template type in golang
import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// define struct that points to template folder location
type Template struct{
	// key templates of pointer type of base template.Template
	templates *template.Template
}
// create struct method, render, for Template type
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	// write data from disk
	// name is template name
	// interface used to write data
	// context

	return t.templates.ExecuteTemplate(w, name, data)

}
// define where to find template folder
func LoadTemplate() *Template {
	// get os path
	path, _ := os.Executable()

	// get file path without the app added to it
	filePath := filepath.Dir(path)

	//template location
	templateFolder := fmt.Sprintf("%v/repository/templates/*", filePath)

	//get path to template folder
	template := &Template{
		templates: template.Must(template.ParseGlob(templateFolder)),
	}

	//return template path
	return template
}