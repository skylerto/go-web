package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-type", "text/html")
		//tmpl, err := template.New("test").Parse(doc)
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		template.New("footer").Parse(footer)

		context := Context{
			[3]string{"Lemon", "Orange", "Apple"},
			"The Title",
		}
		templates.Lookup("test").Execute(w, context)

	})

	http.ListenAndServe(":8000", nil)
}

const header = `<!DOCTYPE html>
<html>
  <head>
    <title>{{.}}</title>
  </head>
`

const doc = `
{{template "header" .Title}}
  <body>
    <h1>List of Fruit</h1>
    <ul>
      {{range .Fruit}}
      <li>{{.}}</li>
      {{end}}
    </ul>
  </body>
  {{template "footer"}}
`

const footer = `</html>`

//Context used as a test to inject text into template
type Context struct {
	Fruit [3]string
	Title string
}
