package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Parsing Data
func ParseFormData(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("Form Data Parsing Error")
	}
}

// Unable to find data
func UnableToGetData(err error) {
	if err != nil {
		log.Println("Unable to find data")
	}
}
 //if insert is not possible then give this message
func UnableToInsertData(err error) {
	if err != nil {
		log.Fatalln("Unable to Insert Data ", err)
	}
}  


// Integer to string conversion
func IntToStringConversion(id int32) string {
	t := strconv.Itoa(int(id))
	return t
}



// Template error check
func UnableToFindHtmlTemplate(tmpl *template.Template) {
	if tmpl == nil {
		log.Println("Unable to find html template")
		return
	}
}
// Exectuion error
func ExcutionTemplateError(err error) {
	if err != nil {
		log.Println("Error executing tempalte : ", err)
		return
	}
}