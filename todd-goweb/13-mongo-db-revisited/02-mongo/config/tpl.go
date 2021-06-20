package config

import "html/template"

// TPL is a template bucket
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
