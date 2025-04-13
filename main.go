package main

import (
	"log/slog"

	"github.com/ltlaitoff/2048/web-server"
)

func main() {
	slog.Info("Start webserver")
	webserver.Start()
}
