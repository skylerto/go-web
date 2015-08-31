package controllers

import (
	"net/http"
	"text/template"

	"github.com/skylerto/main/views"
)

type categoriesController struct {
	template *template.Template
}

func (cc *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := views.GetCategories()

	w.Header().Add("Content-Type", "text/html")
	cc.template.Execute(w, vm)
}
