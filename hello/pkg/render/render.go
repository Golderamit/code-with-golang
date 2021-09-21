package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc ,err := CreateTemplateCache()
	 if err != nil {
		 log.Fatal(err)
	 }
	  t,ok := tc[tmpl]
        if !ok {
			log.Fatal(err)
		}

	 buf := new(bytes.Buffer)

	 _ = t.Execute(buf,nil)

	 _,err =buf.WriteTo(w)
	   if err != nil {
		   fmt.Println("error writing template to browser",err)
	   }
} 
/* CreateTemplateCache create a cache as a map */
func CreateTemplateCache ()  (map[string]*template.Template, error) {
	 myCashe := map[string]*template.Template{}

       pages, err := filepath.Glob("./templates/*.page.html")
	    if  err != nil {
		   return myCashe, err
	    }

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println("page is currently",page)

	  ts, err := template.New(name).Funcs(functions).ParseFiles(page)
	  if  err != nil {
		return myCashe,err
	   }

	matches, err := filepath.Glob("./templates/*.layout.html")
	if  err != nil {
		return myCashe,err
	  }
	if len(matches) > 0 {
		ts, err = ts.ParseGlob("./templates/*.layout.html")
		if  err != nil {
			return myCashe,err
		}
		myCashe[name] = ts
	}
	return myCashe,nil
 }
   return myCashe ,nil
}