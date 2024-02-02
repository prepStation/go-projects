package templates

import (
	"html/template"
	"log"
	"net/http"
)

type Login struct {
	BAlertUser bool
	AlertMsg   string
}

type Register struct {
	BAlertUser bool
	AlertMsg   string
}

type Restricted struct {
	CSRF     string
	AlertMsg string
}

var templates = template.Must(template.ParseFiles("./server/templates/templatefiles/login.tmpl", "./server/templates/templatefiles/register.tmpl", "./server/templates/templatefiles/restricted.tmpl"))

func RenderTemplates(w http.ResponseWriter, tmpl string, p any) {
	err := templates.ExecuteTemplate(w, tmpl+".tmpl", p)
	if err != nil {
		log.Printf("template error here: %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
