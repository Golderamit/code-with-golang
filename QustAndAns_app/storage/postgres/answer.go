package postgres

import (
	"log"
	"QustAndAns_app/storage"
)

const createAnswerQ = `
	INSERT INTO answer
	(question_id, user_id,AnswerDetails) VALUES (:question_id, :user_id, :answer_details)
	RETURNING id, created_at, updated_at
`

func (s *Storage) SaveAnswerdb(answer storage.Answer) (int32, error) {
	stmt, err := s.db.PrepareNamed(createAnswerQ)

	if err != nil {
		log.Println(err)
		return 0, err
	}
	var id int32
	if err := stmt.Get(&id, answer); err != nil {
		return 0, err
	}

	return id, nil

}

const AnswerQuery = `select * from answer where question_id=$1 `

func (s *Storage) AnswerQuery(id int32) ([]storage.Answer, error) {
	answer := make([]storage.Answer, 0)
	if err := s.db.Select(&answer, AnswerQuery, id); err != nil {
		return nil, err
	}
	return answer, nil
}

const createLike = `UPDATE answer SET likes=$1 WHERE id=$2;`

func (s *Storage) Createlike(answer storage.Answer) error {
	_, err := s.db.Exec(createLike, answer.Likes, answer.ID)

	if err != nil {
		return err
	}

	return nil
}



const getAnswerA = `SELECT * FROM answer WHERE id=$1 AND question_id=$2  limit 1`

func (s *Storage) GetAnswerDetail(id, question_id int32) (*storage.Answer, error) {
	answer := storage.Answer{}
	if err := s.db.Get(&answer, getAnswerA, id, question_id); err != nil {
		return nil, err
	}
	return &answer, nil

}
