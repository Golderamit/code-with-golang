package handler

import (
	"html/template"
	"log"
	"net/http"
	"question_answer_app/storage"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
)
type (
	SignUpTempData struct {
		CSRFField  template.HTML
		Form       storage.User
		FormErrors map[string]string
	}
)

func (s *Server) getUserLogin(w http.ResponseWriter, r *http.Request) {

	formData := SignUpTempData{
		CSRFField: csrf.TemplateField(r),
	}

	session, _ := s.session.Get(r, "question_answer_app")
	userId:=session.Values["user_id"] 

	if _,ok:=userId.(string);ok{
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}


	s.SignupTemplate(w, r, formData)

}

func (s *Server) postUserLogin(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	var form storage.User
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatalln("Decode Error")
	}

	if err := form.Validate(); err != nil {
		vErros := map[string]string{}

		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					vErros[key] = value.Error()
				}
			}
		}
		data := SignUpTempData{
			CSRFField:  csrf.TemplateField(r),
			Form:       form,
			FormErrors: vErros,
		}

		s.SignupTemplate(w, r, data)
		return
	}

	id, err := s.store.CreateUser(form)
	if err != nil {
		log.Println("data not saved")
	}

	log.Println(id)

	log.Printf("\n %#v", form)

	http.Redirect(w, r, "/login/", http.StatusTemporaryRedirect)

}

func (s *Server) SignupTemplate(w http.ResponseWriter, r *http.Request, form SignUpTempData) {
	temp := s.templates.Lookup("user_login.html")

	if err := temp.Execute(w, form); err != nil {
		log.Fatalln("executing template: ", err)
		return
	}
}
