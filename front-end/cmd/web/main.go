package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
)

const port = "90"

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles(
		"./cmd/web/templates/test.page.gohtml",
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	)
	if err != nil {
		panic(fmt.Sprintf("Error parsing template: %s", err))
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Starting front end service on port:", port)

	err := http.ListenAndServe(net.JoinHostPort("", port), nil)
	if err != nil {
		log.Panic(err)
	}
}
