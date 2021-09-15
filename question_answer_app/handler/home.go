package handler

import (
	"log"
	"net/http"
)
  
type (
	templateData struct {
		UserLogIn bool
	}
)

func (s *Server) getHome(w http.ResponseWriter, r *http.Request) {
	temp := s.templates.Lookup("home.html")
	if temp == nil {
		log.Println("unable to look home.html")
	
	}

	session, _ := s.session.Get(r, "question_answer_app")
	userID := session.Values["user_id"]
    
	if _, ok := userID.(string); ok {
		data := templateData{
			UserLogIn: true,
		}
		if err := temp.Execute(w, data); err != nil {
			log.Fatalln("Session Execution error")
		}
		return
	}

	data := templateData{
		UserLogIn: false,
	}
	if err := temp.Execute(w, data); err != nil {
		log.Fatalln("temp Execution error")
	}
}