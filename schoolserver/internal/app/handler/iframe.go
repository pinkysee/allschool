package handler

import (
	"net/http"

	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type IframeHandler struct {
	ser *services.Services
}

func IframeHandlerinit(s *services.Services) *IframeHandler {
	return &IframeHandler{ser: s}
}

func (s *IframeHandler) AddIframe(c echo.Context) error {
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
	d := &model.Iframe{Path: c.FormValue("Name"), Page: c.FormValue("Page")}
	err = s.ser.Iframe.AddIframe(d, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(200)
}
func (s *IframeHandler) GetIframe(c echo.Context) error {
	page := c.QueryParam("page")
	iframe, err := s.ser.Iframe.GetIframe(page)
	if err != nil {
		return err
	}
	return c.JSON(200, iframe)
}
func (s *IframeHandler) DeleteIframe(c echo.Context) error {
	d := new(model.Iframe)
	if err := c.Bind(d); err != nil {
		return err
	}
	err := s.ser.Iframe.DeleteIframe(d)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}
func (s *IframeHandler) GetAllIframe(c echo.Context) error {
	iframe, err := s.ser.Iframe.GetAllIframe()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, iframe)
}
