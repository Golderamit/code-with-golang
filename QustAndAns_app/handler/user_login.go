package handler

import (
	"html/template"
	"log"
	"net/http"
	"QustAndAns_app/storage"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
)
type (
	user struct {
	User storage.User
  }
)
type UserFormData struct {
	CSRFField  template.HTML
	Form       storage.User
	FormErrors map[string]string
}

func (s *Server) getCreateUser(w http.ResponseWriter, r *http.Request) {
	tmp := s.templates.Lookup("user-login.html")
	UnableToFindHtmlTemplate(tmp)
	usr, err := s.store.GetUser()
	UnableToGetData(err) 
	tempData := user{
		User: *usr,
	}
	log.Println("Method : Create user called.")
	data := UserFormData{
		CSRFField: csrf.TemplateField(r),
	}
	s.loadUserTemplate(w, r, data)

	session, _ := s.session.Get(r, "QustAndAns_app")
	userId:=session.Values["user_id"] 

	if _,ok:=userId.(string);ok{
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err = tmp.Execute(w, tempData)
	ExcutionTemplateError(err)


}



func (s *Server) postCreateUser(w http.ResponseWriter, r *http.Request) {
	ParseFormData(r)
	var creds storage.User
	if err := s.decoder.Decode(&creds, r.PostForm); err != nil {
		log.Fatalln("Decoding error")
	}
	if err := creds.Validate(); err != nil {
		vErrs := map[string]string{}
		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					vErrs[key] = value.Error()
				}
			}
		}
		data := UserFormData{
			CSRFField:  csrf.TemplateField(r),
			Form:       creds,
			FormErrors: vErrs,
		}
		s.loadUserTemplate(w, r, data)
		return
	}
	pass := creds.Password
	hashed, err := HashAndSalt(pass)
	creds.Password = hashed
	_, err = s.store.CreateUser(creds)
	if err !=nil{
		UnableToInsertData(err)
	}
	
	http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
}
	

func (s *Server) loadUserTemplate(w http.ResponseWriter, r *http.Request, form UserFormData) {
	tmpl := s.templates.Lookup("user-login.html")
	UnableToFindHtmlTemplate(tmpl)
	err := tmpl.Execute(w, form)
	ExcutionTemplateError(err)
}
    

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}
/*--------------------Compare Two Password---------------------*/
func ComparePassword(result *storage.User, form Login, w http.ResponseWriter, r *http.Request) {
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(form.Password)); err != nil {
		log.Println("Password does not match.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
