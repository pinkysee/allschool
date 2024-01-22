package repository

import (
	"fmt"

	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
)

type NewsRepository struct {
	db *sqlx.DB
}

func NewsRepositoryinit(db *sqlx.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

/* func (s *NewsRepository) Createnews(c *model.News) error */
func (s *NewsRepository) GetNews() ([]model.News, error) {
	var newss []model.News
	err := s.db.Select(&newss, "SELECT * FROM  news")
	return newss, err
}

func (s *NewsRepository) DeleteNews(id int) error {
	_, err := s.db.Exec("DELETE FROM news WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *NewsRepository) CreateNews(c *model.News) error {
	var news []model.News
	query := fmt.Sprintf("INSERT INTO news (title, text, preview, created_at) VALUES ($1, $2, $3, $4)")
	row := s.db.QueryRow(query, c.Title, c.Text, c.Preview, c.Created_at)
	if err := row.Scan(&news); err != nil {
		return err
	}
	return nil
}
