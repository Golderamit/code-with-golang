package handler

import (
	"log"
	"net/http"
	"strings"
	"text/template"

)
type (
	templateData struct{
		Name string
		Age int
		Address string
	}
)
func (s *Server) getHome(w http.ResponseWriter, r *http.Request){

	funcMap := template.FuncMap{
		"title": func(name string) string {
			return strings.Title(name)
		},
	}
	tmp, _:= template.New("home.html").Funcs(funcMap).ParseFiles("./assets/templates/home.html")
 
	tmpData := templateData{
		Name:    "amit",
		Age:     28,
		Address: "Kda ,khulna-9000",
		
	}
   err := tmp.Execute(w , tmpData)
   if err !=nil {
	   log.Println("error execute templete:",err)
	   return
   }
}