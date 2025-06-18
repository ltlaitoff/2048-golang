package webserver

import (
	"net/http"

	"github.com/ltlaitoff/2048/auth"
	"github.com/ltlaitoff/2048/core"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.IsAuthorizedSession(r)

	InitialRender(w, session)
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.IsAuthorizedSession(r)

	Render(w, session)
}

func keyHandler(action string, w http.ResponseWriter, r *http.Request) {
	session, _ := auth.IsAuthorizedSession(r)

	core.Move(action, session, "Test agent")
	Render(w, session)
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
	session, _ := auth.IsAuthorizedSession(r)
	Render(w, session)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.AuthSignUp(w, r)
	InitialRender(w, session)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.AuthSignIn(w, r)
	InitialRender(w, session)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.IsAuthorizedSession(r)

	core.Reset(session)
	Render(w, session)
}
