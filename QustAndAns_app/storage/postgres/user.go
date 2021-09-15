package postgres

 import (
	"QustAndAns_app/storage"
 )

 const getUser = `
	SELECT * from users
	WHERE id = $1 and password=$2
	
  `

  func (s *Storage) GetUser() (*storage.User, error) {
	user := storage.User{}
	if err := s.db.Get(&user, getUser); err != nil {
		return nil, err
	}
	return &user, nil
   }

  const createUser = `
	INSERT INTO users(
		first_name,
		last_name,
		username,
		email,
		password
	)
	VALUES(
		:first_name,
		:last_name,
		:username,
		:email,
		:password
	)
	RETURNING id,created_at, updated_at
	`

	func (s *Storage) CreateUser(usr storage.User) (int32, error) {
		stmt, err := s.db.PrepareNamed(createUser)
		if err != nil {
			return 0, err
		}
		var id int32
		if err := stmt.Get(&id, usr); err != nil {
			return 0, err
		}
		return id, nil
	}

	const userQ = `SELECT * FROM users WHERE email=$1`

	func (s *Storage) GetUserInfo(email string) *storage.User {
		user := storage.User{}
		err := s.db.Get(&user, userQ, email)
		if err != nil {
			return &user
		}
		return &user
	}
	
	const getUserID = `SELECT id FROM users WHERE email=$1 LIMIT 1`

    func (s *Storage) GetUserdb(email interface{}) (int32, error) {
		user := storage.User{}
	
		if _, ok := email.(string); ok {
			if err := s.db.Get(&user, getUserID, email); err != nil {
				return 0, err
			}
		}
	
		return user.ID, nil
	}