package controllers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/skylerto/main/controllers/util"
	"github.com/skylerto/main/views"
)

type standLocatorController struct {
	template *template.Template
}

func (this *standLocatorController) get(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := views.GetStandLocator()
	responseWriter.Header().Add("Content Type", "text/html")
	this.template.Execute(responseWriter, vm)
}

func (this *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	//vm := views.GetStandLocations()
	vm := views.GetStandLocator()
	responseWriter.Header().Add("Content Type", "application/json")
	data, err := json.Marshal(vm)
	if err == nil {
		responseWriter.Write(data)
	} else {
		responseWriter.WriteHeader(404)
	}
}
