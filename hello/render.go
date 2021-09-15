package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate,_ :=template.ParseFiles("./assets/templates/"+tmpl)
	err :=parsedTemplate.Execute(w,nil)
	if err != nil{
		fmt.Println("error when parsing template ",err)
	}
  return 
} 