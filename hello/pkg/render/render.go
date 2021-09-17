package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"google.golang.org/genproto/googleapis/cloud/functions/v1"
)

var functions = template.FuncMap{

}

func RenderTemplate(w http.ResponseWriter, tmpl string){
	_,err :=RenderTemplateRun(w)
	  if err != nil{
		  fmt.Println("error in getting template cache:",err)
	  }
	ParsedTemplate,_ :=template.ParseFiles("./assets/templates/"+tmpl)
	err :=ParsedTemplate.Execute(w,nil)
	if err != nil{
		fmt.Println("error when parsing template ",err)
	}
  return 
} 

func RenderTemplateRun(w http.ResponseWriter) (map[string]*template.Template, error){
	myCashe := map[string]*template.Template{}

      pages, err := filepath.Glob("./templates/*.page.html")
	    if  err != nil {
		   return myCashe,err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println("page is currently",page)

	  ts, err :=template.New(name).Funcs(functions).ParseFiles(page)
	  if  err != nil {
		return myCashe,err
	}

	matches, err :=filepath.Glob("./templates/*.layout.html")
	if  err != nil {
		return myCashe,err
	}
	if lens(matches) > 0{
		ts, err := ts.ParseGlob("./templates/*.layout.html")
		if  err != nil {
			return myCashe,err
		}
		myCashe[name] = ts
	}
	return myCashe,nil
}