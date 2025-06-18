package webserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ltlaitoff/2048/auth"
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

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read body
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user auth.SignUpUserBody
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = auth.SignUpUser(user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	core.Reset()
	Render(w)
}
