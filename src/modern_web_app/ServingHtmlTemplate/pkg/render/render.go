package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmbl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmbl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]

	if !inMap {
		//need to create template
	} else {
		//we have the template in the cache
		log.Println("using cached template")
	}

}
