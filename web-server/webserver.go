package webserver

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"

	"github.com/ltlaitoff/2048/core"
	"github.com/ltlaitoff/2048/db"
	"github.com/ltlaitoff/2048/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const DEBUG = true

//go:embed assets/*
var assetsFiles embed.FS

func Start() {
	database, err := db.ConnectMongoDb("mongodb://root:example@localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	var results []entities.User

	collection := database.Database("2048").Collection("users")
	finded, err2 := collection.Find(context.Background(), bson.D{})

	if err2 != nil {
		log.Fatal(err2)
	}

	defer finded.Close(context.Background())

	err3 := finded.All(context.Background(), &results)

	if err3 != nil {
		log.Fatal(err3)
	}

	// finded.Decode(&results)
	log.Println(results)

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
