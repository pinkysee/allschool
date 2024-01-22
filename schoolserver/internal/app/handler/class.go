package handler

import (
	"net/http"

	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	ser *services.Services
}

func ClassHandlerinit(s *services.Services) *ClassHandler {
	return &ClassHandler{ser: s}
}
func (s *ClassHandler) AddClass(c echo.Context) error {
	class := new(model.Class)
	if err := c.Bind(class); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err := s.ser.Class.AddClass(class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(200)
}
func (s *ClassHandler) GetClass(c echo.Context) error {
	class, err := s.ser.Class.GetClass()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(200, class)
}
func (s *ClassHandler) DeleteClass(c echo.Context) error {
	class := new(model.Class)
	if err := c.Bind(class); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err := s.ser.Class.DeleteClass(class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
