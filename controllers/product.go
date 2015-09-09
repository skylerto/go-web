package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/skylerto/main/controllers/util"
	"github.com/skylerto/main/views"
)

type product struct {
	template *template.Template
}

func (pc *product) get(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := views.GetProduct(id)
		w.Header().Add("Content-Type", "text/html")
		pc.template.Execute(w, vm)
	} else {
		w.WriteHeader(404)
	}
}

type productController struct {
	template *template.Template
}

func (cc *productController) get(w http.ResponseWriter, req *http.Request) {
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
