package render

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Erro on getting template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*page.tmpl")
	if err != nil {
		return myCache, nil
	}

	for _, page := range pages {
		fName := filepath.Base(page)

		ts, err := template.New(fName).ParseFiles(page)
		if err != nil {
			return myCache, nil
		}

		matches, err := filepath.Glob("../../templates/*layout.tmpl")
		if err != nil {
			return myCache, nil
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*layout.tmpl")
			if err != nil {
				return myCache, nil
			}
		}

		myCache[fName] = ts
	}
	return myCache, nil
}
