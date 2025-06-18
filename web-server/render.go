package webserver

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ltlaitoff/2048/core"
	"github.com/ltlaitoff/2048/entities"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

type RenderData struct {
	Score           core.Score
	Cells           core.Board
	IsEnd           bool
	IsAuthenticated bool
}

var tempatePaths []string

func InitRender(statisFiles embed.FS) {
	entries, err := fs.ReadDir(assetsFiles, "assets/templates")

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		tempatePaths = append(tempatePaths, "assets/templates/"+entry.Name())
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

		b, err := fs.ReadFile(assetsFiles, filename)
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

func InitialRender(w http.ResponseWriter, session *entities.Session) {
	drivers := template.Must(compileTemplates(tempatePaths...))
	driver, err := drivers.Clone()

	if err != nil {
		log.Fatal("Cloning helpers: ", err)
	}

	cells, score, end := core.State(session)

	var data RenderData

	data.Cells = cells
	data.Score = score
	data.IsEnd = end
	data.IsAuthenticated = session != nil

	driver.ExecuteTemplate(w, "index.html", data)
}

func Render(w http.ResponseWriter, session *entities.Session) {
	if session == nil {
		InitialRender(w, session)
		return
	}

	drivers := template.Must(compileTemplates(tempatePaths...))
	driver, err := drivers.Clone()

	if err != nil {
		log.Fatal("Cloning helpers: ", err)
	}

	cells, score, end := core.State(session)

	var data RenderData

	data.Cells = cells
	data.Score = score
	data.IsEnd = end

	driver.ExecuteTemplate(w, "Root", data)
}
