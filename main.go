package main

import (
	"bufio" // buffered output
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/", new(MyHandler))

	// MUX: Creates a new thread
	http.ListenAndServe(":8000", nil)
}

// SayHi just responds
func SayHi(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is a test of function passing"))
}

//MyHandler struct wrapper for http.Handler
type MyHandler struct {
	http.Handler
}

//ServeHTTP Serves up a file
func (handler *MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	f, err := os.Open(path) // buffered output

	//data, err := ioutil.ReadFile(string(path))

	if err == nil {
		bufferedReader := bufio.NewReader(f)
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}

		res.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(res) // buffered output
		//	res.Write(data)
	} else {
		res.WriteHeader(404)
		res.Write([]byte("404 - " + http.StatusText(404)))
	}
}
