package server

import (
	"fmt"
	"kresimir.dev/modules/templater"
	"log"
	"net/http"
)

type Server struct {
	port int
}

func CreateServer(port int) *Server {
	return &Server{
		port: port,
	}
}

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func MakeHandleFunc(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Fatal(err)
		}
	}
}

func (server *Server) Listen() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", MakeHandleFunc(server.handleIndex))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", server.port), nil), nil)
}

func (server *Server) handleIndex(w http.ResponseWriter, r *http.Request) error {
	log.Println("Processing a handle index request.")
	return templater.GenerateTemplate(w, 200, "/index.html", nil)
}