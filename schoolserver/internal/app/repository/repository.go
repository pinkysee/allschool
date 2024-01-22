package repository

import (
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type User interface {
	DeleteUserByClass(s *model.User) error
	GetUser(s *model.User) (model.User, error)
	CheckUser(s *model.User) (model.User, error)
	AddUser(s *model.User) (int, error)
	DeleteUser(id int) (string, error)
	GetAllTeacher() ([]model.User, error)
	GetAllZav() ([]model.User, error)
	GetUserByClass(s *model.User) ([]model.User, error)
}
type News interface {
	GetNews() ([]model.News, error)
	DeleteNews(id int) error
	CreateNews(c *model.News) error
}
type Docs interface {
	GetDocs(page string) ([]model.Docs, error)
	GetAllDocs() ([]model.Docs, error)
	AddDocs(c *model.Docs) (string, error)
	DeleteDocs(path string) (string, error)
}
type Test interface {
	CreateTest(c *model.Test) (int, error)
	DeleteTest(c *model.Test) error
	GetTests(name string, classname string) ([]model.Test, error)
	GetTestAnswer(id int) ([]model.Test, error)
	GetTestQuetion(id int) ([]model.Test, error)
	SetAnswer(c *model.Test) error
	GetRightQuetion(id int) ([]model.Test, error)
}
type Rasp interface {
	GetRasp(c *model.User) ([]model.Rasp, error)
	SetRasp(c *model.Rasp) error
	GetAllRasp() ([]model.Rasp, error)
}
type Iframe interface {
	AddIframe(c *model.Iframe) error
	DeleteIframe(c *model.Iframe) error
	GetIframe(page string) ([]model.Iframe, error)
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
type Repository struct {
	User
	News
	Docs
	Test
	Rasp
	Iframe
	Notification
	Class
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:         UserRepositoryinit(db),
		News:         NewsRepositoryinit(db),
		Docs:         DocsRepositoryinit(db),
		Test:         TestRepositoryinit(db),
		Rasp:         RaspRepositoryinit(db),
		Iframe:       IframeRepositoryInit(db),
		Notification: NotifRepositoryInit(db),
		Class:        ClassRepositoryinit(db),
	}
}
