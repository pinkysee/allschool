package handler

import (
	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

// zav

type User interface {
	GetUserByClass(c echo.Context) error
	DeleteUserByClass(c echo.Context) error
	Auth(c echo.Context) error
	Register(c echo.Context) error
	GetAllTeachers(c echo.Context) error
	GetAllZav(c echo.Context) error
	Delete(c echo.Context) error
	CheckJWT(c echo.Context) error
}
type News interface {
	GetNews(c echo.Context) error
	DeleteNews(c echo.Context) error
	CreateNews(c echo.Context) error
}
type uvd interface {
	GetNotification(c echo.Context) error
	CreateNotification(c echo.Context) error
}
type test interface {
	Settest(c *echo.Context) error
	Gettest(c *echo.Context) error
	Setanswertest(c *echo.Context) error
}
type rasp interface {
	GetRasp(c echo.Context) error
	SetRasp(c echo.Context) error
	GetAllRasp(c echo.Context) error
}
type docs interface {
	Getdocs(c echo.Context) error
	Adddocs(c echo.Context) error
	DeleteDocs(c echo.Context) error
	GetAllDocs(c echo.Context) error
}
type Iframe interface {
	AddIframe(c echo.Context) error
	GetIframe(c echo.Context) error
	DeleteIframe(c echo.Context) error
	GetAllIframe(c echo.Context) error
}
type Class interface {
	AddClass(c echo.Context) error
	GetClass(c echo.Context) error
	DeleteClass(c echo.Context) error
}
type Handling struct {
	User
	News
	uvd
	test
	rasp
	docs
	Iframe
	Class
}

func Handlinginit(s *services.Services) *Handling {
	return &Handling{
		User:   Userhandlerinit(s),
		News:   Newshandlerinit(s),
		uvd:    Uvdhandlerinit(s),
		test:   Testhandlerinit(s),
		rasp:   rasphandlerinit(s),
		docs:   Docshandlerinit(s),
		Iframe: IframeHandlerinit(s),
		Class:  ClassHandlerinit(s),
	}
}

type er struct {
	Message string
}
