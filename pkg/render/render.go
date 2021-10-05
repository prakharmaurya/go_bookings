package render

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/prakharmaurya/go_bookings/pkg/config"
	"github.com/prakharmaurya/go_bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(rw http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var t *template.Template
	var ok bool
	if app.UseCache {
		t, ok = app.TemplateCache[tmpl]
	} else {
		tc, err := CreateTemplateCache()
		if err != nil {
			fmt.Println("Failed to create template cache")
		}
		t, ok = tc[tmpl]
	}

	if !ok {
		fmt.Println("failed to get template by string name", ok)
		return
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)

	if err != nil {
		fmt.Println("Error in executing template", err)
	}
	_, err = buf.WriteTo(rw)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
