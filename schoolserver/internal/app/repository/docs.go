package repository

import (
	"fmt"

	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DocsRepository struct {
	db *sqlx.DB
}

func DocsRepositoryinit(db *sqlx.DB) *DocsRepository {
	return &DocsRepository{db: db}
}

func (s *DocsRepository) GetDocs(page string) ([]model.Docs, error) {
	var docs []model.Docs
	err := s.db.Select(&docs, "SELECT * FROM docs WHERE page = $1", page)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return docs, nil
}
func (s *DocsRepository) AddDocs(c *model.Docs) (string, error) {
	var name string
	query := fmt.Sprintf("INSERT INTO docs (name, path, page) VALUES ($1, $2, $3) RETURNING name")
	row := s.db.QueryRow(query, c.Name, c.Path, c.Page)
	if err := row.Scan(&name); err != nil {
		logrus.Error(err)
		return "", err
	}
	return name, nil
}

func (s *DocsRepository) DeleteDocs(path string) (string, error) {
	_, err := s.db.Exec("DELETE FROM docs WHERE path = $1", path)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return "sucefull delete docs from db", nil
}
func (s *DocsRepository) GetAllDocs() ([]model.Docs, error) {
	var docs []model.Docs
	err := s.db.Select(&docs, "SELECT * FROM docs")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return docs, nil
}
