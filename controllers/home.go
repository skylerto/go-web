package controllers

import (
	"net/http"
	"text/template"

	"github.com/skylerto/main/controllers/util"
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
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	home.template.Execute(responseWriter, vm)
}
