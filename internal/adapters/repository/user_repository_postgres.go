package repository

import (
	"database/sql"

	"crudWithDB/internal/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

func (r *UserRepositoryPostgres) Create(u domain.User) error {
	_, err := r.db.Exec(
		`INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`,
		u.ID, u.Name, u.Email,
	)
	return err
}

func (r *UserRepositoryPostgres) GetByID(id string) (domain.User, error) {
	var u domain.User
	err := r.db.Get(&u, `SELECT id, name, email FROM users WHERE id=$1`, id)
	return u, err
}

func (r *UserRepositoryPostgres) Update(u domain.User) error {
	res, err := r.db.Exec(
		`UPDATE users SET name=$1, email=$2 WHERE id=$3`,
		u.Name, u.Email, u.ID,
	)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *UserRepositoryPostgres) Delete(id string) error {
	res, err := r.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
