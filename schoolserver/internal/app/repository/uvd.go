package repository

import (
	"fmt"

	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type NotifRepository struct {
	db *sqlx.DB
}

func NotifRepositoryInit(db *sqlx.DB) *NotifRepository {
	return &NotifRepository{db: db}
}
func (s *NotifRepository) CreateNotification(c *model.Uvd) error {
	query := fmt.Sprintf("INSERT INTO notification (Text, Class) VALUES ($1, $2)")
	_, err := s.db.Exec(query, c.Text, c.Class)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (s *NotifRepository) GetNotification(c *model.User) ([]model.Uvd, error) {
	var notif []model.Uvd
	err := s.db.Select(&notif, `SELECT * FROM  notification WHERE Class = $1`, c.Classname)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return notif, nil
}
