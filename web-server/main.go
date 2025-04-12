package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"ltlaitoff/2048/core"
)

var assets string = ""

func render(w http.ResponseWriter) {
	cells := ""

	core.Map(func (value int64) {
			cellValue := strconv.FormatInt(value, 10)

			if cellValue == "0" {
				cellValue = ""
			}

			cells += fmt.Sprintf("<cell data-value=\"%s\">%s</cell>", cellValue, cellValue)
	})

	fmt.Fprintf(w, cells)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(assets + "/assets/templates/index.html")

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	render(w)
}

func upHandler(w http.ResponseWriter, r *http.Request) {
	core.Up()
	render(w)
}

func leftHandler(w http.ResponseWriter, r *http.Request) {
	core.Left()
	render(w)
}

func rightHandler(w http.ResponseWriter, r *http.Request) {
	core.Right()
	render(w)
}

func downHandler(w http.ResponseWriter, r *http.Request) {
	core.Down()
	render(w)
}

func enterHandler(w http.ResponseWriter, r *http.Request) {
	render(w)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	core.Reset()
	render(w)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf(dir)

	assets = dir + "/web-server"

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/clicked", clickedHandler)

	http.HandleFunc("/enter", enterHandler)
	http.HandleFunc("/reset", resetHandler)

	http.HandleFunc("/up", upHandler)
	http.HandleFunc("/left", leftHandler)
	http.HandleFunc("/right", rightHandler)
	http.HandleFunc("/down", downHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assets+"/assets"))))

	log.Printf("Started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
