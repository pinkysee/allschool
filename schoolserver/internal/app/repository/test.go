package repository

import (
	"fmt"

	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
)

type TestRepository struct {
	db *sqlx.DB
}

func TestRepositoryinit(db *sqlx.DB) *TestRepository {
	return &TestRepository{db: db}
}

func (s *TestRepository) CreateTest(c *model.Test) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO test (name, lesson, komu, deadline, answer, quetion, teachername, rightans) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id")
	row := s.db.QueryRow(query, c.Name, c.Lesson, c.Komu, c.Deadline, c.Answer, c.Quetion, c.Teachername, c.Rightans)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TestRepository) DeleteTest(c *model.Test) error {
	_, err := s.db.Exec("DELETE FROM test WHERE id = $1", c.Id)
	if err != nil {
		return err
	}
	return nil
}
func (s *TestRepository) GetTests(name string, classname string) ([]model.Test, error) {
	var test []model.Test
	err := s.db.Select(&test, "SELECT id, komu, name, lesson, deadline FROM	test WHERE komu = $1 OR kome LIKE %$2% ", classname, name)
	return test, err
}
func (s *TestRepository) GetTestAnswer(id int) ([]model.Test, error) {
	var test []model.Test
	err := s.db.Select(&test, "SELECT answer FROM test WHERE id=$1 ", id)
	return test, err
}
func (s *TestRepository) GetTestQuetion(id int) ([]model.Test, error) {
	var test []model.Test
	err := s.db.Select(&test, "SELECT quetion FROM test WHERE id=$1 ", id)
	return test, err
}
func (s *TestRepository) SetAnswer(c *model.Test) error {
	_, err := s.db.Exec("UPDATE your_table SET ff = CONCAT(ff, $1) WHERE id = $2", c.Answer, c.Id)
	if err != nil {
		return err
	}
	return nil
}
func (s *TestRepository) GetRightQuetion(id int) ([]model.Test, error) {
	var test []model.Test
	err := s.db.Select(&test, "SELECT rightans FROM test WHERE id=$1 ", id)
	return test, err
}

/* func (s *Repository) setanswertest(c *model.Test) error */
