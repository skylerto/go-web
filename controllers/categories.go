package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/skylerto/main/controllers/util"
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

type categoryController struct {
	template *template.Template
}

func (cc *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)

	if err == nil {
		vm := views.GetProducts(id)
		w.Header().Add("Content-type", "text/html")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()

		cc.template.Execute(responseWriter, vm)

	} else {
		w.WriteHeader(404)
	}
}
