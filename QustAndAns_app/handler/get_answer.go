package handler

import (

	"html/template"
	"log"
	"net/http"
	"QustAndAns_app/storage"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type (
	AnswerTempData struct {
		QuesionDetails storage.Question
		CSRFField      template.HTML
		Form           storage.Answer
		FormErrors     map[string]string
	}
)

func (s *Server) getAnswer(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	questionId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
	}

	question, err := s.store.GetQuestionDetail(int32(questionId))

	if err != nil {
		log.Fatalln(err)
	}

	formData := AnswerTempData{
		QuesionDetails: *question,
		CSRFField:      csrf.TemplateField(r),
	}

	s.AnswerCreateTemplate(w, r, formData)

}

func (s *Server) postAnswer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	questionId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
	}

	question, err := s.store.GetQuestionDetail(int32(questionId))

	if err != nil {
		log.Fatalln(err)
	}
	if err := r.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	var form storage.Answer
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatalln("Decode Error")
	}

	if err := form.ValidateAnswer(); err != nil {
		vErros := map[string]string{}

		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					vErros[key] = value.Error()
				}
			}
		}
		data := AnswerTempData{
			QuesionDetails: *question,
			CSRFField:      csrf.TemplateField(r),
			Form:           form,
			FormErrors:     vErros,
		}

		s.AnswerCreateTemplate(w, r, data)
		return
	}

	session, _ := s.session.Get(r, "QustAndAns_app")
	user_email := session.Values["user_email"]

	user_id, err := s.store.GetUserdb(user_email)

	if err != nil {
		log.Fatalln(err)
	}

	form.UserID = user_id
	form.QuestionID = question.ID
	
	log.Println(form)

	a, err:= s.store.SaveAnswerdb(form)
	
	if  err!= nil {
		log.Println("data not saved")
	}

	log.Println(a)



	http.Redirect(w, r, "/get/question", http.StatusSeeOther)

}


func (s *Server) AnswerCreateTemplate(w http.ResponseWriter, r *http.Request, form AnswerTempData) {
	temp := s.templates.Lookup("create-answer.html")

	if err := temp.Execute(w, form); err != nil {
		log.Fatalln("executing template: ", err)
		return
	}
}
