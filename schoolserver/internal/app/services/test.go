package services

import (
	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/model"
)

type TestRepository struct {
	rep *repository.Repository
}

func TestRepositoryinit(rep *repository.Repository) *TestRepository {
	return &TestRepository{rep: rep}
}
func (s *TestRepository) CreateTest(c *model.Test) error { return nil }
func (s *TestRepository) DeleteTest(c *model.Test) error { return nil }
func (s *TestRepository) SetAnswer(c *model.Test) error  { return nil }
func (s *TestRepository) GetTests(name string, classname string) ([]model.Test, error) {
	return s.rep.Test.GetTests(name, classname)
}
func (s *TestRepository) GetTestAnswer(id int) ([]model.Test, error)   { return nil, nil }
func (s *TestRepository) GetRightQuetion(id int) ([]model.Test, error) { return nil, nil }

func (s *TestRepository) GetTestQuetion(id int) ([]model.Test, error) { return nil, nil }
