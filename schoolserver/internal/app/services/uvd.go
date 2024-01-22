package services

import (
	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/model"
)

type UvdRepository struct {
	rep *repository.Repository
}

func UvdRepositoryinit(rep *repository.Repository) *UvdRepository {
	return &UvdRepository{rep: rep}
}

func (s *UvdRepository) CreateNotification(c *model.Uvd) error {
	return s.rep.Notification.CreateNotification(c)
}
func (s *UvdRepository) GetNotification(c *model.User) ([]model.Uvd, error) {
	return s.rep.Notification.GetNotification(c)
}
