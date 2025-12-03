package handlers

import "html/template"

//Used Ref: go.dev to figure how to put tmpl and Init in one place
// so we don’t have to duplicate them across all handler files:

// Shared template instance for all handlers
var tmpl *template.Template

// Init sets the template from main.go
func Init(t *template.Template) {
	tmpl = t
}
