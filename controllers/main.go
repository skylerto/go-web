package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

//Register registers the controllers
func Register(templates *template.Template) {
	router := mux.NewRouter()

	/*
	*	Setup a new route
	 */
	//create a new pointer to the homeController
	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	router.HandleFunc("/home", hc.get)

	/*
	*	Setup a new route
	 */
	//create a new pointer to the homeController
	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	router.HandleFunc("/categories", cc.get)

	categoryController := new(categoryController)
	categoryController.template = templates.Lookup("products.html")
	router.HandleFunc("/categories/{id}", categoryController.get)

	productController := new(productController)
	productController.template = templates.Lookup("product.html")
	router.HandleFunc("/products/{id}", productController.get)

	profileController := new(profileController)
	profileController.template = templates.Lookup("profile.html")
	router.HandleFunc("/profile", profileController.handle)

	standLocatorController := new(standLocatorController)
	standLocatorController.template = templates.Lookup("stand_locator.html")
	router.HandleFunc("/stand_locator", standLocatorController.get)
	router.HandleFunc("/api/stand_locator", standLocatorController.apiSearch)

	http.Handle("/", router)

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
