package server

import (
	"fmt"
	"kresimir.dev/modules/data"
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
			http.Redirect(w, r, "/404", http.StatusSeeOther)
		}
	}
}

func (server *Server) Listen() {
	fs := http.FileServer(http.Dir("./static"))
	templater.ParseTemplates()

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// TODO: this router is hacky, especially 404 handling
	http.HandleFunc("/", MakeHandleFunc(server.handleIndex))
	http.HandleFunc("/post/", MakeHandleFunc(server.handlePost))
	http.HandleFunc("/404", MakeHandleFunc(server.handleNotFound))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", server.port), nil), nil)
}

func (server *Server) handleNotFound(w http.ResponseWriter, r *http.Request) error {
	log.Println("Processing get /404.html request.")
	return templater.GenerateTemplate(w, "404", nil)
}

func (server *Server) handleIndex(w http.ResponseWriter, r *http.Request) error {
	// TODO: this 404 handling doesn't belong here
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	}
	log.Println("Processing get /index.html request.")
	return templater.GenerateTemplate(w, "index", data.Posts)
}

// TODO: rewrite this, each post shouldn't be a template
func (server *Server) handlePost(w http.ResponseWriter, r *http.Request) error {
	slug := r.URL.Path[len("/post/"):]
	log.Println("Processing get", slug, "request.")
	return templater.GenerateTemplate(w, slug, nil)
}
