package repository

import (
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type IframeRepository struct {
	db *sqlx.DB
}

func IframeRepositoryInit(db *sqlx.DB) *IframeRepository {
	return &IframeRepository{db: db}
}
func (s *IframeRepository) AddIframe(c *model.Iframe) error {
	query := "INSERT INTO iframe (page, path) VALUES ($1, $2)"
	_, err := s.db.Exec(query, c.Page, c.Path)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
func (s *IframeRepository) DeleteIframe(c *model.Iframe) error {
	_, err := s.db.Exec("DELETE FROM iframe WHERE path = $1", c.Path)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
func (s *IframeRepository) GetIframe(page string) ([]model.Iframe, error) {
	var Iframe []model.Iframe
	err := s.db.Select(&Iframe, "SELECT * FROM iframe WHERE page = $1", page)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return Iframe, nil
}
func (s *IframeRepository) GetAllIframe() ([]model.Iframe, error) {
	var Iframe []model.Iframe
	err := s.db.Select(&Iframe, "SELECT * FROM iframe")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return Iframe, nil
}
