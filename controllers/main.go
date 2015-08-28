package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/skylerto/main/views"
)

func Register(templates *template.Template) {
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template :=
				templates.Lookup(requestedFile + ".html")

			var context interface{} = nil
			switch requestedFile {
			case "home":
				context = views.GetHome()
			case "categories":
				context = views.GetCategories()
			case "products":
				context = views.GetProducts()
			case "product":
				context = views.GetProduct()
			}
			if template != nil {
				template.Execute(w, context)
			} else {
				w.WriteHeader(404)
			}
		})

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
