package handler

import (
	"net/http"

	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Docshandler struct {
	ser *services.Services
}

func Docshandlerinit(s *services.Services) *Docshandler {
	return &Docshandler{ser: s}
}
func (s *Docshandler) Getdocs(c echo.Context) error {
	d := new(model.Docs)
	if err := c.Bind(d); err != nil {
		c.NoContent(http.StatusBadRequest)
	}
	docs, err := s.ser.Docs.GetDocs(d.Page)
	if err != nil {
		return err
	}

	logrus.Info(docs)
	return c.JSON(200, docs)
}
func (s *Docshandler) Adddocs(c echo.Context) error {
	file, err := c.FormFile("File")
	if err != nil {
		logrus.Error(err)
		jsontoken := &Api{Token: "Bad Request"}
		return c.JSON(http.StatusBadRequest, jsontoken)
	}
	if c.FormValue("Name") == "" || c.FormValue("Page") == "" {
		jsontoken := &Api{Token: "Bad Request"}
		return c.JSON(http.StatusBadRequest, jsontoken)
	}
	d := &model.Docs{Name: c.FormValue("Name"), Page: c.FormValue("Page")}
	name, err := s.ser.Docs.AddDocs(d, file)
	if err != nil || name == "" {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, name)
}

func (s *Docshandler) DeleteDocs(c echo.Context) error {
	d := new(model.Docs)
	if err := c.Bind(d); err != nil {
		return err
	}
	a, err := s.ser.Docs.DeleteDocs(d.Path)
	if err != nil {
		return err
	}
	return c.JSON(200, a)
}
func (s *Docshandler) GetAllDocs(c echo.Context) error {
	docs, err := s.ser.Docs.GetAllDocs()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, docs)
}
