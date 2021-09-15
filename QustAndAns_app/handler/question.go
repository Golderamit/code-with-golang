package handler

import (

	"html/template"
	"log"
	"net/http"
	"QustAndAns_app/storage"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/csrf"
)

 

type QuestionFormTempData struct {
	CSRFField  template.HTML
	Form       storage.Question
	FormErrors map[string]string
}

func (s *Server) getCreateQuestion(w http.ResponseWriter, r *http.Request) {
	log.Println("method: createQuestion")
	data := QuestionFormTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       storage.Question{},
		FormErrors: map[string]string{},
	}
	s.loadCreateQuestionTemplate(w, r, data)
}

/*func (s *Server) getQuestionShow(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
	fmt.Fprint(w, "Hello from question Details: ", id)
}  */

func (s *Server) postCreateQuestion(w http.ResponseWriter, r *http.Request) {
	log.Println("method: postQuestion")
	if err := r.ParseForm(); err != nil {
		log.Fatalln("parsing error")
	}

	var form storage.Question
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatalln("decoding error")
	}

	if err:= form.Validate(); err != nil {
		vErrs := map[string]string{}
		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value:= range e {
					vErrs[key] = value.Error()
				}
			}
		}
		
		data := QuestionFormTempData{
			CSRFField: csrf.TemplateField(r),
			Form: form,
			FormErrors: vErrs,
		}
		s.loadCreateQuestionTemplate(w, r, data)
		return
	} 
	session, _ := s.session.Get(r, "QustAndAns_app")
	user_email := session.Values["user_email"]

	user_id, err := s.store.GetUserdb(user_email)


	if err != nil {
		log.Fatalln(err)
	}
	
	form.UserID = user_id
	

	id, err := s.store.CreateQuestion(form)
	if err != nil {
		log.Fatalln("unable to save data")
	}
	log.Println(id)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func (s *Server) loadCreateQuestionTemplate(w http.ResponseWriter, r *http.Request, form QuestionFormTempData) {
	tmp := s.templates.Lookup("ask-question.html")
	if err := tmp.Execute(w, form); err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
