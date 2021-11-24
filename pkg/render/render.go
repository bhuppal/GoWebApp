package render

import (
	"bytes"
	"fmt"
	"github.com/bhuppal/go/goweb/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc := app.TemplateCache

	/*
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
*/
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
	/*
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error occured during parsing template file ", err)
	}
	*/

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
			return myCache, nil
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