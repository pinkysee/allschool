package services

import (
	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/model"
)

type Classervices struct {
	rep *repository.Repository
}

func Classervicesinit(repo *repository.Repository) *Classervices {
	return &Classervices{rep: repo}
}

func (s *Classervices) AddClass(c *model.Class) error {
	return s.rep.Class.AddClass(c)
}
func (s *Classervices) GetClass() ([]model.Class, error) {
	return s.rep.Class.GetClass()
}
func (s *Classervices) DeleteClass(c *model.Class) error {
	return s.rep.Class.DeleteClass(c)
}
