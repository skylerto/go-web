package controllers

import (
	"net/http"
	"text/template"

	"github.com/skylerto/main/views"
)

//homeController
type homeController struct {
	template *template.Template
}

//get binds the writer
func (home *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := views.GetHome()

	w.Header().Add("Content-type", "text/html")
	home.template.Execute(w, vm)
}
