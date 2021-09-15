package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string){
	ParsedTemplate,_ :=template.ParseFiles("./assets/templates/"+tmpl)
	err :=ParsedTemplate.Execute(w,nil)
	if err != nil{
		fmt.Println("error when parsing template ",err)
	}
  return 
} 