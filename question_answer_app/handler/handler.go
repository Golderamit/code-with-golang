package handler

import (
	"question_answer_app/storage/postgres"
	"html/template"
	"net/http"
     "log"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
)

type (
 Server struct {
  templates   *template.Template
  store       *postgres.Storage
  decoder     *schema.Decoder
  session      *sessions.CookieStore
 }
)

func NewServer(st *postgres.Storage,decoder *schema.Decoder,session *sessions.CookieStore) (*mux.Router, error) {

  s := &Server{
	 templates: &template.Template{},
 	 store:     st,
	 decoder: decoder,
	 session: session,
 }
   if err := s.parseTemplates(); err != nil {
	log.Println("parse template error")
	   
}
    r := mux.NewRouter()
	r.Use(csrf.Protect([]byte("1234")))

	 /*csrf.Protect([]byte("it-is-secret------"), csrf.Secure(false))(r) */

	/* staic files Handler */
	/*r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))*/
	
	r.HandleFunc("/", s.getHome)

    /* Login Handler */
	r.HandleFunc("/login/", s.getLogin).Methods("GET")
	r.HandleFunc("/login/", s.postLogin).Methods("POST")
    r.HandleFunc("/logout/", s.logout).Methods("GET")

	/* User Create Handler */
	r.HandleFunc("/user/login/", s.getUserLogin).Methods("GET")
	r.HandleFunc("/user/login/", s.postUserLogin).Methods("POST")
	
	
	/*... user question handler....*/
	ar := r.NewRoute().Subrouter()
	ar.Use(s.authMiddleware)
	
	ar.HandleFunc("/question/create/", s.getCreateQuestion).Methods("GET")
	ar.HandleFunc("/question/create/", s.postCreateQuestion).Methods("POST")


    ar.HandleFunc("/get/question/", s.getQuestionList).Methods("GET")
	ar.HandleFunc("/get/question/{id}/", s.createQuestionDetails).Methods("POST")

	ar.HandleFunc("/question/create/answer/{id}/", s.getAnswer).Methods("GET")
	ar.HandleFunc("/question/create/answer/{id}/", s.postAnswer).Methods("POST")
	
       return r, nil
}

func (s *Server) parseTemplates() error {
	templates := template.New("templates").Funcs(template.FuncMap{
	 "strrev": func(str string) string {
	  n := len(str)
	  runes := make([]rune, n)
	  for _, rune := range str {
	   n--
	   runes[n] = rune
	  }
	  return string(runes[n:])
	 },
	}).Funcs(sprig.FuncMap())
   
	tmpl, err := templates.ParseGlob("assets/templates/*.html")
	if err != nil {
	 return err
	}
	s.templates = tmpl
	return nil
   }

   func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.session.Get(r, "question_answer_app_session")
		value := session.Values["user_id"]
		user_email := session.Values["user_email"]
		log.Println(value)
		if _, ok := user_email.(string); ok {
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden,you r nt registerd", http.StatusForbidden)
		}
	})
}

func (s *Server) loadCreateQuestionTemplate(w http.ResponseWriter, r *http.Request, form QuestionFormTempData) {
	tmp := s.templates.Lookup("ask-question.html")
	if err := tmp.Execute(w, form); err != nil {
		log.Println("Error executing template :", err)
		return
	}
}