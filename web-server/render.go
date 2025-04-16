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
	IsEnd bool
}

var assetsPath string = ""
var tempatePaths []string

func InitRender(path string) {
	assetsPath = path
	tempatePaths = []string{
		assetsPath + "/assets/templates/index.html",
		assetsPath + "/assets/templates/root.html",
		assetsPath + "/assets/templates/cell.html",
		assetsPath + "/assets/templates/end.html",
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

	cells, score, end := core.State()

	var data RenderData

	data.Cells = cells
	data.Score = score
	data.IsEnd = end

	driver.ExecuteTemplate(w, "index.html", data)
}

func Render(w http.ResponseWriter) {
	drivers := template.Must(compileTemplates(tempatePaths...))
	driver, err := drivers.Clone()

	if err != nil {
		log.Fatal("Cloning helpers: ", err)
	}

	cells, score, end := core.State()

	var data RenderData

	data.Cells = cells
	data.Score = score
	data.IsEnd = end

	driver.ExecuteTemplate(w, "Root", data)
}

