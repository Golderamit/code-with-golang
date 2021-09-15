package handler

import (
	"net/http"
)

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	session, _ := s.session.Get(r, "question_answer_app")
	session.Values["user_id"] = nil
	session.Values["is_admin"] = nil

	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}