package templates

import (
	"html/template"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	Templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.Templates.ExecuteTemplate(w, name, data)
}

func ParseTemplates(path string) (*template.Template, error) {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
		"Round":   Round,
	})
	err := filepath.Walk(path, func(path string, _ os.FileInfo, err error) error {
		if strings.Contains(path, ".gohtml") {
			_, err = t.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return t, nil
}

func Round(value float64) float64 {
	return math.Round(value*100) / 100
}
