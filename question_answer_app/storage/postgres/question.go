package postgres

import (
	"question_answer_app/storage"
)

const createQuestion = `
INSERT INTO question 
(user_id, QuestionTitle) 
VALUES (:user_id, :QuestionTitle) 
RETURNING id, created_at, updated_at
`

const getQuestion = `SELECT * FROM question `
const getAQuestion = `SELECT * FROM question WHERE id=$1 limit 1`

func (s *Storage) CreateQuestion(question storage.Question) (int32, error) {
	stmt, err := s.db.PrepareNamed(createQuestion)
	if err != nil {
		return 0, err
	}

	var id int32
	if err := stmt.Get(&id, question); err != nil {
		return 0, err
	}
	return id, nil
}
func (s *Storage) GetQuestionListDB() ([]storage.Question, error) {
	question := make([]storage.Question, 0)
	if err := s.db.Select(&question, getQuestion); err != nil {
		return nil, err
	}

	return question, nil
}
func (s *Storage) GetQuestionDetail(id int32) (*storage.Question, error) {
	question := storage.Question{}
	if err := s.db.Get(&question, getAQuestion, id); err != nil {
		return nil, err
	}
	return &question, nil

}