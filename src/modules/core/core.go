package core

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	port int
}

type ServerHandlerFunc func(http.ResponseWriter, *http.Request) error

func CreateServer(port int) *Server {
	return &Server{
		port: port,
	}
}

func (server *Server) Listen() {
	//TODO: Extract to router?
	http.HandleFunc("/", MakeServerHandleFunc(server.handleIndex))
	http.ListenAndServe(fmt.Sprintf("%s%d", ":", server.port), nil)
}

func MakeServerHandleFunc(f ServerHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Fatal(err)
		}
	}
}

// this belongs to a templater

func GenerateTemplate(w http.ResponseWriter, status int, name string, data any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html")
	template := template.Must(template.ParseFiles(fmt.Sprintf("%s%s", "./templates", name)))
	return template.Execute(w, data)
}

// Routes

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return GenerateTemplate(w, 200, "/index.html", nil)
}
