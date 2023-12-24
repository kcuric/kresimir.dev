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

		defer func() {
			if rec := recover(); rec != nil {
				fmt.Println("Fatal error:", r)
				http.Redirect(w, r, "/404", http.StatusSeeOther)
			}
		}()

		if err := f(w, r); err != nil {
			log.Fatal(err)
		}
	}
}

func (server *Server) Listen() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// TODO: this router is hacky, especially 404 handling
	http.HandleFunc("/", MakeHandleFunc(server.handleIndex))
	http.HandleFunc("/post/", MakeHandleFunc(server.handlePost))
	http.HandleFunc("/404", MakeHandleFunc(server.handleNotFound))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", server.port), nil), nil)
}

func (server *Server) handleNotFound(w http.ResponseWriter, r *http.Request) error {
	log.Println("Processing get /404.html request.")
	return templater.GenerateTemplate(w, "/404.html", nil)
}

func (server *Server) handleIndex(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	}
	log.Println("Processing get /index.html request.")
	return templater.GenerateTemplate(w, "/index.html", data.Posts)
}

func (server *Server) handlePost(w http.ResponseWriter, r *http.Request) error {
	slug := r.URL.Path[len("/post"):]
	name := slug + ".html"
	log.Println("Processing get", name, "request.")
	return templater.GenerateTemplate(w, name, nil)
}
