package repositories

import (
	"database/sql"
	"cat-social/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	query := "SELECT id, email, name, password FROM users WHERE email = $1 AND deleted_at IS NULL"
	row := r.db.QueryRow(query, email)

	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Create(user *entities.User) error {
	query := "INSERT INTO users (id, email, name, password) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, user.ID, user.Email, user.Name, user.Password)
	return err
}
