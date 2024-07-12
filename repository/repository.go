package repository

import (
	"github.com/Anu-renjith/gin-sqlx/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	query := "INSERT INTO users (name, email) VALUES (:name, :email)"
	result, err := r.db.NamedExec(query, &user)
	if err != nil {
		return user, err
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	// Set the ID of the user
	user.ID = uint(id)
	return user, nil
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Select(&users, "SELECT id, name, email FROM users")
	return users, err
}
