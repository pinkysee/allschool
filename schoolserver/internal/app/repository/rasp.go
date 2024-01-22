package repository

import (
	"fmt"

	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RaspRepository struct {
	db *sqlx.DB
}

func RaspRepositoryinit(db *sqlx.DB) *RaspRepository {
	return &RaspRepository{db: db}
}

func (s *RaspRepository) GetRasp(c *model.User) ([]model.Rasp, error) {
	var rasp []model.Rasp
	err := s.db.Select(&rasp, "SELECT * FROM schedule WHERE classname = $1", c.Classname)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return rasp, nil
}

func (s *RaspRepository) SetRasp(c *model.Rasp) error {
	query := fmt.Sprintf("INSERT INTO schedule (lessons, classname) VALUES ($1, $2)")
	_, err := s.db.Exec(query, c.Lessons, c.Classname)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (s *RaspRepository) GetAllRasp() ([]model.Rasp, error) {
	var rasp []model.Rasp
	err := s.db.Select(&rasp, "SELECT * FROM schedule")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return rasp, nil
}
