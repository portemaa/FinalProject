package handlers

import (
    "html/template"
    "net/http"
)

// can be extended if we needed logging 
type AboutHandler struct{}

func NewAboutHandler() *AboutHandler {
    return &AboutHandler{}
}

func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("assets/about.html"))
    tmpl.Execute(w, nil)
}
