package handler

import (
	"html/template"
	"net/http"

	"class_2/practice/storage/postgres"
	"github.com/Masterminds/sprig"
	"github.com/gorilla/mux"
)

type (
	Server struct {
		templates *template.Template
		store     *postgres.Storage
	}
)

func NewServer(st *postgres.Storage) (*mux.Router, error) {
	s := &Server{
		store: st,
	}

	if err := s.parseTemplates(); err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/", s.getHome)
	/* r.HandleFunc("/articles", s.getArticleList).Methods("GET")
	r.HandleFunc("/articles/{id}", s.getArticleShow).Methods("GET") */

	return r, nil
}

func (s *Server) parseTemplates() error {
	templates := template.New("templates").Funcs(template.FuncMap{
		"strrev": func(str string) string {
			n := len(str)
			runes := make([]rune, n)
			for _, rune := range str {
				n--
				runes[n] = rune
			}
			return string(runes[n:])
		},
	}).Funcs(sprig.FuncMap())

	tmpl, err := templates.ParseGlob("assets/templates/*.html")
	if err != nil {
		return err
	}
	s.templates = tmpl
	return nil
}
