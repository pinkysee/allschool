package services

import (
	"mime/multipart"

	"github.com/PINKYSEE/schoolserver/internal/app/repository"

	"github.com/PINKYSEE/schoolserver/internal/model"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type User interface {
	GetUserByClass(c *model.User) ([]model.User, error)
	DeleteUserByClass(c *model.User) error
	GenerateJWT(c *model.User) (string, error)
	Register(c *model.User) (int, error, string)
	Delete(userID int) (string, error)
	GetAllTeacher() ([]model.User, error)
	GetAllZav() ([]model.User, error)
	CheckJWT(jw string) (*model.User, error)
}
type News interface {
	GetNews() ([]model.News, error)
	CreateNews(c *model.News, v *multipart.FileHeader) error
	DeleteNews(c *model.News) error
}
type Docs interface {
	GetDocs(page string) ([]model.Docs, error)
	DeleteDocs(name string) (string, error)
	GetAllDocs() ([]model.Docs, error)
	AddDocs(c *model.Docs, v *multipart.FileHeader) (string, error)
}
type Rasp interface {
	GetRasp(c *model.User) (*DaysRasp, error)
	SetRasp(c *model.Rasp) error
	GetAllRasp() ([]*AllRasp, error)
}
type Iframe interface {
	AddIframe(c *model.Iframe, v *multipart.FileHeader) error
	GetIframe(page string) ([]model.Iframe, error)
	DeleteIframe(c *model.Iframe) error
	GetAllIframe() ([]model.Iframe, error)
}
type Notification interface {
	CreateNotification(c *model.Uvd) error
	GetNotification(c *model.User) ([]model.Uvd, error)
}
type Class interface {
	AddClass(c *model.Class) error
	GetClass() ([]model.Class, error)
	DeleteClass(c *model.Class) error
}
type Services struct {
	User
	News
	Docs
	Rasp
	Iframe
	Notification
	Class
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{
		User:         UserServiceinit(rep),
		News:         NewsServiceinit(rep),
		Docs:         DocsRepositoryinit(rep),
		Rasp:         RaspRepositoryinit(rep),
		Iframe:       IframeServicesinit(rep),
		Notification: UvdRepositoryinit(rep),
		Class:        Classervicesinit(rep),
	}
}
