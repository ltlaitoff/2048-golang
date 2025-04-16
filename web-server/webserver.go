package webserver

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ltlaitoff/2048/core"
)

const DEBUG = true

func Start() {
	core.Init()

	if DEBUG {
		slog.SetLogLoggerLevel(slog.LevelDebug.Level())
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	assets := dir + "/web-server"
	InitRender(assets)

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/clicked", clickedHandler)

	http.HandleFunc("/enter", enterHandler)
	http.HandleFunc("/reset", resetHandler)

	http.HandleFunc("/top", topHandler)
	http.HandleFunc("/left", leftHandler)
	http.HandleFunc("/right", rightHandler)
	http.HandleFunc("/bottom", bottomHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assets+"/assets"))))

	slog.Info("Started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
