package repository

import (
	"fmt"

	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ClassRepository struct {
	db *sqlx.DB
}

func ClassRepositoryinit(db *sqlx.DB) *ClassRepository {
	return &ClassRepository{db: db}
}
func (s *ClassRepository) AddClass(c *model.Class) error {
	query := fmt.Sprintf("INSERT INTO class (Name) VALUES ($1)")
	_, err := s.db.Exec(query, c.Name)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
func (s *ClassRepository) GetClass() ([]model.Class, error) {
	var Class []model.Class
	err := s.db.Select(&Class, `SELECT * FROM  class`)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return Class, nil
}
func (s *ClassRepository) DeleteClass(c *model.Class) error {
	_, err := s.db.Exec("DELETE FROM class WHERE Name = $1", c.Name)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
