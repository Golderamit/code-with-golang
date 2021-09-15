package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) getArticleList(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello from Article List")
}

func (s *Server) getArticleShow(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
	fmt.Fprint(w, "Hello from Article Details: ", id)
}