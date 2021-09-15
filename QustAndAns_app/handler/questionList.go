package handler

import (

	"html/template"
	"log"
	"net/http"
	"QustAndAns_app/storage"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type (
	TemplateData struct {
		Question []storage.Question
	}

	TempDataDetails struct {
		CSRFField template.HTML
		Question  storage.Question
		Answer   []storage.Answer
		UserId    int32
	}
)

func (s *Server) getQuestionList(w http.ResponseWriter, r *http.Request) {

	questionDb, err := s.store.GetQuestionListDB()
	if err != nil {
		log.Println("database error")
	}
	temp := TemplateData{
		Question: questionDb,
		
	}
    

	s.GetQuestionListTemplate(w, r,"question_list.html",temp)
}

func (s *Server) createQuestionDetails(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	questionId, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	question, err := s.store.GetQuestionDetail(int32(questionId))
	if err != nil {
		log.Println("do not match")
	}

	answerSet, err := s.store.AnswerQuery(question.ID)

	if err != nil {
		log.Fatalln(err)
	}

	session, _ := s.session.Get(r, "QustAndAns_app")
	user_email := session.Values["user_email"]

	user_id, err := s.store.GetUserdb(user_email)

	if err != nil {
		log.Fatalln(err)
	}

	temp := TempDataDetails{
		CSRFField: csrf.TemplateField(r),
		Question:  *question,
		Answer:    answerSet,
		UserId:    user_id,
	}

	s.GetQuestionListTemplate(w, r, "question_list.html",temp)
}




func (s *Server) GetQuestionListTemplate(w http.ResponseWriter, r *http.Request, temp_name string, data interface{}) {
	temp := s.templates.Lookup(temp_name)

	if err := temp.Execute(w, data); err != nil {
		log.Println("Error executing template :", err)
		return
	}
}