package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if err == nil {
			context := Context{
				[3]string{"Lemon", "Orange", "Apple"},
				"The Title",
			}
			tmpl.Execute(w, context)
		}
	})

	http.ListenAndServe(":8000", nil)
}

const doc = `<!DOCTYPE html>
<html>
  <head>
    <title>{{.Title}}</title>
  </head>
  <body>
    <h1>List of Fruit</h1>
    <ul>
      {{range .Fruit}}
      <li>{{.}}</li>
      {{end}}
    </ul>
  </body>
</html>`

//Context used as a test to inject text into template
type Context struct {
	Fruit [3]string
	Title string
}
