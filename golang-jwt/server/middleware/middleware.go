package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"golang-jwt/db"
	"golang-jwt/server/middleware/token"
	"golang-jwt/server/templates"

	"github.com/justinas/alice"
)

func NewHandler() http.Handler {
	return alice.New(recoverHandler, authHandler).ThenFunc(logicHandler)
}

func logicHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/restricted":
		csrfSecret := grabCsrfFromRequest(r)
		templates.RenderTemplates(w, "restricted", &templates.Restricted{CSRF: csrfSecret, AlertMsg: "Hello Akhil"})
	case "/login":
		switch r.Method {
		case "GET":
		case "POST":

		default:

		}
	case "/register":
		switch r.Method {
		case "GET":
			templates.RenderTemplates(w, "register", &templates.Register{BAlertUser: false, AlertMsg: ""})
		case "POST":

			r.ParseForm()
			log.Println(r.Form)

			_, uuid, err := db.FetchUserByUsername(strings.Join(r.Form["username"], ""))
			if err == nil {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				role := "user"
				uuid, err := db.StoreUser(strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""), role)
				if err != nil {
					http.Error(w, http.StatusText(500), 500)
				}
				log.Printf("uuid: %s", uuid)

				authToken, refreshTOken, csrfSecret, err := token.CreateNewToken(uuid, role)
				if err != nil {
					http.
				}
			}
		default:

		}
	case "/logout":
	case "/deleteuser":
	default:
	}

}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Panic("Recovered: panic:%+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func authHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/restricted", "/logout", "/deleteuser":
		default:
		}

	}
}

func nullifyTokenCookies(w http.ResponseWriter, r *http.Request) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &authCookie)
	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &refreshCookie)
	refreshTokenCookie, err := r.Cookie("RefreshToken")
	if err == http.ErrNoCookie {
		return
	} else if err != nil {
		log.Panicf("panic:%+v", err)
		http.Error(w, http.StatusText(500), 500)
	}
	token.RevokeRefreshToken(refreshTokenCookie.Value)
}

func setAuthAndRefreshCookies(w http.ResponseWriter, authToken, refreshToken string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authToken,
		HttpOnly: true,
	}

	http.SetCookie(w, &authCookie)
	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshToken,
		HttpOnly: true,
	}

	http.SetCookie(w, &refreshCookie)

}

func grabCsrfFromRequest(r *http.Request) string {

	csrfToken := r.FormValue("X-CSRF-Token")
	if csrfToken == "" {
		return csrfToken
	}
	return r.Header.Get("X-CSRF-Token")
}
