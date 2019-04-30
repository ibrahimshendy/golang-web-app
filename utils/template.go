package utils

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func LoadTemplate(target string) {
	tpl = template.Must(template.ParseGlob(target))
}

func LoadView(w http.ResponseWriter, tmp string, data interface{}) {
	tpl.ExecuteTemplate(w, tmp, data)
}
