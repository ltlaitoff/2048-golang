package webserver

import (
	"log"
	"net/http"

	"github.com/ltlaitoff/2048/auth"
	"github.com/ltlaitoff/2048/core"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated, err := auth.IsAuthorized(r)

	if err != nil {
		log.Fatal(err)
	}

	InitialRender(w, *isAuthenticated)
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated, err := auth.IsAuthorized(r)

	if err != nil {
		log.Fatal(err)
	}
	Render(w, *isAuthenticated)
}

func keyHandler(action string, w http.ResponseWriter, r *http.Request) {
	isAuthenticated, err := auth.IsAuthorized(r)

	if err != nil {
		log.Fatal(err)
	}

	core.Move(action)

	Render(w, *isAuthenticated)
}

func topHandler(w http.ResponseWriter, r *http.Request) {
	keyHandler("TOP", w, r)
}

func leftHandler(w http.ResponseWriter, r *http.Request) {
	keyHandler("LEFT", w, r)
}

func rightHandler(w http.ResponseWriter, r *http.Request) {
	keyHandler("RIGHT", w, r)
}

func bottomHandler(w http.ResponseWriter, r *http.Request) {
	keyHandler("BOTTOM", w, r)
}

func enterHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated, err := auth.IsAuthorized(r)

	if err != nil {
		log.Fatal(err)
	}

	Render(w, *isAuthenticated)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	auth.AuthSignUp(w, r)

	InitialRender(w, true)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	auth.AuthSignIn(w, r)

	InitialRender(w, true)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated, err := auth.IsAuthorized(r)

	if err != nil {
		log.Fatal(err)
	}

	core.Reset()
	Render(w, *isAuthenticated)
}
