package repository

import (
	model "cat-social/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	//IsEmailUsed(email string) (bool, error)
	EmailIsExist(email string) bool
}
type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	_, err := r.db.Exec(context.Background(), "INSERT INTO users (email, name, password) VALUES ($1, $2, $3)", user.Email, user.Name, user.Password)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(context.Background(), "SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) EmailIsExist(email string) bool {
	var exist string
	err := r.db.QueryRow(context.Background(), "SELECT email FROM users WHERE email = $1 LIMIT 1", email).Scan(&exist)
	fmt.Println(exist)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
	}
	return true
}
