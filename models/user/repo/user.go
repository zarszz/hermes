package repo

import (
	"hermes/controllers/entity"
	models "hermes/models/user"

	"github.com/jmoiron/sqlx"
)

type UserService struct {
	connection *sqlx.DB
}

func NewUserService (conn *sqlx.DB) *UserService {
	return &UserService{connection: conn}
}

func (s *UserService) Create(user entity.UserRegistrationInput) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"

	tx := s.connection.MustBegin()
	tx.MustExec(query, user.Name, user.Email, user.Password)
	err := tx.Commit()
	
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	user := []models.User{}
	err := s.connection.Select(&user, query, email)
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, nil
	}
	return &user[0], nil
}

func (s *UserService) GetAll() (*models.Users, error) {
	query := `SELECT * FROM users`
	var users models.Users
	err := s.connection.Select(&users.Users, query)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (s *UserService) GetByPK(id int) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	user := []models.User{}
	err := s.connection.Select(&user, query, id)
	if err != nil {
		return nil, err
	}
	if len(user) == 0{
		return nil, nil
	}
	return &user[0], nil
}

func (s *UserService) UpdateOne(user entity.UserUpdateInput, id int) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`

	tx := s.connection.MustBegin()
	tx.MustExec(query, user.Name, user.Email, id)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteOne(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	tx := s.connection.MustBegin()
	tx.MustExec(query, id)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
