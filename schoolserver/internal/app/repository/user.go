package repository

import (
	"fmt"

	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *sqlx.DB
}

func UserRepositoryinit(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) GetUser(s *model.User) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE login = $1", s.Login)
	return user, err
}
func (r *UserRepository) GetUserByClass(s *model.User) ([]model.User, error) {
	var user []model.User
	err := r.db.Select(&user, "SELECT * FROM users WHERE classname = $1", s.Classname)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) DeleteUserByClass(s *model.User) error {
	_, err := r.db.Exec("DELETE FROM users WHERE classname = $1", s.Classname)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) CheckUser(s *model.User) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT password, name, role FROM users WHERE password = $1", s.Login)
	return user, err
}

func (r *UserRepository) AddUser(s *model.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO users (name, login, password, role, classname, avatar) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	row := r.db.QueryRow(query, s.Name, s.Login, s.Password, s.Role, s.Classname, s.Avatar)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) DeleteUser(id int) (string, error) {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (r *UserRepository) GetAllTeacher() ([]model.User, error) {
	var teacher []model.User
	err := r.db.Select(&teacher, `SELECT * FROM  users WHERE role = 'teacher'`)
	if err != nil {
		return nil, err
	}
	return teacher, err
}

func (r *UserRepository) GetAllZav() ([]model.User, error) {
	var zav []model.User
	err := r.db.Select(&zav, `SELECT * FROM  users WHERE role = 'zav'`)
	if err != nil {
		return nil, err
	}
	return zav, err
}
