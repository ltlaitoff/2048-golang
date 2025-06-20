package webserver

import (
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"

	"github.com/ltlaitoff/2048/core"
	"github.com/ltlaitoff/2048/db"
)

const DEBUG = true

//go:embed assets/*
var assetsFiles embed.FS

func Start() {
	err := db.ConnectMongoDb("mongodb://root:example@localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	core.Init()

	if DEBUG {
		slog.SetLogLoggerLevel(slog.LevelDebug.Level())
	}

	InitRender(assetsFiles)

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/clicked", clickedHandler)

	http.HandleFunc("/reset", resetHandler)

	http.HandleFunc("/top", topHandler)
	http.HandleFunc("/left", leftHandler)
	http.HandleFunc("/right", rightHandler)
	http.HandleFunc("/bottom", bottomHandler)

	http.HandleFunc("/auth/sign-up", signUpHandler)
	http.HandleFunc("/auth/sign-in", signInHandler)

	strippedFS, _ := fs.Sub(assetsFiles, "assets")
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.FS(strippedFS))))

	slog.Info("Started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
