package webserver

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ltlaitoff/2048/core"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

type RenderData struct {
	Score core.Score
	Cells core.Board
}

var assetsPath string = ""
var tempatePaths []string

func InitRender(path string) {
	assetsPath = path
	tempatePaths = []string{
		assetsPath + "/assets/templates/index.html",
		assetsPath + "/assets/templates/root.html",
		assetsPath + "/assets/templates/cell.html",
	}
}

func compileTemplates(filenames ...string) (*template.Template, error) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	var tmpl *template.Template
	for _, filename := range filenames {
		name := filepath.Base(filename)
		if tmpl == nil {
			tmpl = template.New(name)
		} else {
			tmpl = tmpl.New(name)
		}

		b, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		mb, err := m.Bytes("text/html", b)
		if err != nil {
			return nil, err
		}
		tmpl.Parse(string(mb))
	}
	return tmpl, nil
}

func InitialRender(w http.ResponseWriter) {
	drivers := template.Must(compileTemplates(tempatePaths...))
	driver, err := drivers.Clone()

	if err != nil {
		log.Fatal("Cloning helpers: ", err)
	}

	cells, score := core.State()

	var data RenderData

	data.Cells = cells
	data.Score = score

	driver.ExecuteTemplate(w, "index.html", data)
}

func Render(w http.ResponseWriter) {
	drivers := template.Must(compileTemplates(tempatePaths...))
	driver, err := drivers.Clone()

	if err != nil {
		log.Fatal("Cloning helpers: ", err)
	}

	cells, score := core.State()

	var data RenderData

	data.Cells = cells
	data.Score = score

	driver.ExecuteTemplate(w, "Root", data)
}

func RenderEnd(w http.ResponseWriter) {
	Render(w)

	// cells := ""
	//
	// core.Map(func(value int64) {
	// 	cellValue := strconv.FormatInt(value, 10)
	//
	// 	if cellValue == "0" {
	// 		cellValue = ""
	// 	}
	//
	// 	cells += fmt.Sprintf("<cell data-value=\"%s\">%s</cell>", cellValue, cellValue)
	// })
	//
	// cells += "<div>Game end xdd!</div>"
	//
	// fmt.Fprintf(w, cells)

	// t, err := template.ParseFiles(assets + "/assets/templates/index.html")
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// t.Execute(w, nil)
}
