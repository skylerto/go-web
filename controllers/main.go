package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"
)

//Register registers the controllers
func Register(templates *template.Template) {
	/*
	*	Setup a new route
	 */
	//create a new pointer to the homeController
	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	http.HandleFunc("/", hc.get)

	/*
	*	Setup a new route
	 */
	//create a new pointer to the homeController
	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	http.HandleFunc("/categories", cc.get)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)

}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
