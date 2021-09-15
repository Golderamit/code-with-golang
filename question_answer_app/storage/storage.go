package storage

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	User struct {
		ID        int32     `db:"id"`
		FirstName string    `db:"first_name"`
		LastName  string    `db:"last_name"`
		Username  string    `db:"username"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		IsActive  bool      `db:"is_active"`
		IsAdmin   bool      `db:"is_admin"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
    Question struct {
		ID            int32      `db:"id"`
		UserID        int32    `db:"user_id"`
		QuestionTitle string     `db:"question_title"`
		CreatedAt    time.Time   `db:"created_at"`
		UpdatedAt    time.Time   `db:"updated_at"`
	}
	Answer struct {
		ID          int32      `db:"id"`
		QuestionID  int32      `db:"question_id"`
		UserID      int32      `db:"user_id"`
		AnswerDetails string   `db:"answer_details"`
		Likes        int32     `db:"likes"`
		CreatedAt  time.Time   `db:"created_at"`
		UpdatedAt  time.Time   `db:"updated_at"`
	}
	
)

const nameLength = "Name should be 4 to 30 Characters"
const usernameLength = "User Name should be 4 to 30 Characters"
const emailLength = "Email should be 4 to 30 Characters"
const passLength = "Password length should be 6 to 30"
const firstNameRequired = "First Name is Required"
const lastNameRequired = "Last Name is Required"
const emailIsRequired = "Email is Required"
const passwordIsRequired = "password is Required"
const usernameIsRequired = "User name is Required"

func (sg User) Validate() error {
	return validation.ValidateStruct(&sg,
		validation.Field(&sg.FirstName,
			validation.Required.Error(firstNameRequired),
			validation.Length(4, 30).Error(nameLength),
		),

		validation.Field(&sg.LastName,
			validation.Required.Error(lastNameRequired),
			validation.Length(4, 30).Error(nameLength),
		),

		validation.Field(&sg.Email,
			validation.Required.Error(emailIsRequired),
			validation.Length(4, 30).Error(emailLength),
		),

		validation.Field(&sg.Username,
			validation.Required.Error(usernameIsRequired),
			validation.Length(4, 30).Error(usernameLength),
		),
		validation.Field(&sg.Password,
			validation.Required.Error(passwordIsRequired),
			validation.Length(6, 30).Error(passLength),
		),
	)
}

func (sg User) ValidateUser() error {
	return validation.ValidateStruct(&sg,
		validation.Field(&sg.Email,
			validation.Required.Error("email is required"),
			is.Email,
		),
		validation.Field(&sg.Password,
			validation.Required.Error("Password is required"),
			validation.Length(3, 10).Error("Password Lenght must be 3 to 10"),
		),
	)
}
func (q Question) Validate() error {
	return validation.ValidateStruct(&q,
		validation.Field(&q.UserID, 
			validation.Required.Error("The  question title field is required."),
			validation.Length(5, 200).Error("The title should be between 5 to 200 characters"),
		),
		validation.Field(&q.QuestionTitle, 
			validation.Required.Error("The question details field is required."),
			validation.Length(10, 1000).Error("The title should be between 10 to 1000 characters"),
		),
	)
}

func (s Answer) ValidateAnswer() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.AnswerDetails,
			validation.Required.Error("Answer is required"),
			validation.Length(2, 350).Error("Answer text length must be 2 to 350"),
		),
	)
}