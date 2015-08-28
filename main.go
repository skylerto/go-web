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
			//context := Context{"Meri"}
			tmpl.Execute(w, req.URL.Path)
		}
	})

	http.ListenAndServe(":8000", nil)
}

const doc = `<!DOCTYPE html>
<html>
  <head>
    <title>Example title</title>
  </head>
  <body>
    {{if eq . "/Google"}}
      <h1>Hello Google</h1>
    {{else}}
      Hello, {{.}}
    {{end}}
  </body>
</html>`

//Context used as a test to inject text into template
type Context struct {
	Message string
}
