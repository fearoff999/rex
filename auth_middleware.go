package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	godotenv.Load(".env.rex")
	user := os.Getenv("REX_USER")
	password := os.Getenv("REX_PASSWORD")

	return func(rw http.ResponseWriter, rq *http.Request) {
		u, p, ok := rq.BasicAuth()
		if !ok || len(strings.TrimSpace(u)) < 1 || len(strings.TrimSpace(p)) < 1 {
			unauthorised(rw)
			return
		}

		if u != user || p != password {
			unauthorised(rw)
			return
		}

		handler(rw, rq)
	}
}

func unauthorised(rw http.ResponseWriter) {
	rw.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	rw.WriteHeader(http.StatusUnauthorized)
}
