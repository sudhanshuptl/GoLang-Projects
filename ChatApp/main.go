// Goal is to to create A simple web server

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Need to write logic to get template and render it
type templateHandler struct {
	once sync.Once   // it insure it fetch th template just once, regardless how many go-routins are calling
					// it is very helpful because web server in go is automatically concurrent
	filename string
	templ  *template.Template
}

// ServeHTTP handles the HTTP request, Name should be exactly same
// this method is binded with templateHandler structure
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// t is return value
	//if once.Do(f) is called multiple times, only the first call will invoke f,
	// even if f has a different value in each invocation. A new instance of Once is required for each function to execute
	t.once.Do(func() {
		pwd, _ := os.Getwd()
		fmt.Println("current file path", pwd) // printed current working directory
		t.templ = template.Must(template.ParseFiles(filepath.Join("ChatApp/templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}


// package name and main function name should be same
func main()  {
	// using own handler
	http.Handle("/", &templateHandler{filename: "home.html"})
	// Start Web server
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("ListenAndServer", err)
	}
}



