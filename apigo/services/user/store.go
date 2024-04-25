package user

import (
	"database/sql"
	"fmt"

	"github.com/codezeron/apigo/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error){
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	users := new(types.User)
	for rows.Next(){
		users, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if users.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return users, nil
}

//verifica o conteudo das linhas trazidas da tabela
func scanRowIntoUser(rows *sql.Rows) (*types.User, error){
	u := new(types.User)
	err := rows.Scan(
		&u.ID,
		&u.FirstName, 
		&u.LastName, 
		&u.Email, 
		&u.CreatedAt,
)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error){
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error{
	return nil
}
