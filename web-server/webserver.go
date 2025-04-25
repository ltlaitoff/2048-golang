package webserver

import (
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"

	"github.com/ltlaitoff/2048/core"
)

const DEBUG = true

//go:embed assets/*
var assetsFiles embed.FS

func Start() {
	core.Init()

	if DEBUG {
		slog.SetLogLoggerLevel(slog.LevelDebug.Level())
	}

	InitRender(assetsFiles)

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/clicked", clickedHandler)

	http.HandleFunc("/enter", enterHandler)
	http.HandleFunc("/reset", resetHandler)

	http.HandleFunc("/top", topHandler)
	http.HandleFunc("/left", leftHandler)
	http.HandleFunc("/right", rightHandler)
	http.HandleFunc("/bottom", bottomHandler)


	strippedFS, _ := fs.Sub(assetsFiles, "assets")
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.FS(strippedFS))))

	slog.Info("Started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
