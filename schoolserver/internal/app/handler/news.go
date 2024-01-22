package handler

import (
	"net/http"

	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Newshandler struct {
	ser *services.Services
}

func Newshandlerinit(s *services.Services) *Newshandler {
	return &Newshandler{ser: s}
}
func (s *Newshandler) GetNews(c echo.Context) error {
	cd, err := s.ser.GetNews()
	if err != nil {
		logrus.Warning(err)
		return c.JSON(http.StatusBadGateway, &er{Message: "Bad Gateway"})
	}
	return c.JSON(200, cd)
}
func (s *Newshandler) DeleteNews(c echo.Context) error {
	news := new(model.News)
	if err := c.Bind(news); err != nil {
		return err
	}
	if news.Preview == "" {

		return c.JSON(http.StatusBadGateway, &er{Message: "Bad Preview"})
	}
	s.ser.News.DeleteNews(news)
	return nil
}
func (s *Newshandler) CreateNews(c echo.Context) error {
	file, err := c.FormFile("File")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if c.FormValue("Title") == "" || c.FormValue("Text") == "" {
		jsontoken := &Api{Token: "Bad Request"}
		return c.JSON(http.StatusBadRequest, jsontoken)
	}
	if err != nil {
		logrus.Info(err)
	}
	d := &model.News{Title: c.FormValue("Title"), Text: c.FormValue("Text")}
	err = s.ser.News.CreateNews(d, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "OK")
}
