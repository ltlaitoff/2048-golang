package webserver

import (
	"net/http"

	"github.com/ltlaitoff/2048/core"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	InitialRender(w)
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	Render(w)
}

func keyHandler(action string, w http.ResponseWriter) {
	core.Move(action)

	Render(w)
}

func topHandler(w http.ResponseWriter, r *http.Request) {
	keyHandler("TOP", w)
}

func leftHandler(w http.ResponseWriter, r *http.Request) {
	keyHandler("LEFT", w)
}

func rightHandler(w http.ResponseWriter, r *http.Request) {
	keyHandler("RIGHT", w)
}

func bottomHandler(w http.ResponseWriter, r *http.Request) {
	keyHandler("BOTTOM", w)
}

func enterHandler(w http.ResponseWriter, r *http.Request) {
	Render(w)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	core.Reset()
	Render(w)
}
