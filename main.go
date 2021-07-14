package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func healthHandler(rw http.ResponseWriter, rq *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func handler(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != "GET" || rq.RequestURI != "/" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	cmd := os.Getenv("REX_COMMAND")
	cmdSplit := strings.Split(cmd, " ")
	out, errCommand := exec.Command(cmdSplit[0], cmdSplit[1:]...).CombinedOutput()
	if errCommand != nil {
		log.Printf("Failed with error %s", errCommand)
	}
	rw.Write(out)
	rw.WriteHeader(http.StatusOK)
}

func RecoverWrap(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("UNKNOWN_ERROR")
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", RecoverWrap(BasicAuth(handler)))
	http.HandleFunc("/health", RecoverWrap(BasicAuth(healthHandler)))

	godotenv.Load(".env.rex")
	port := os.Getenv("REX_PORT")

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
