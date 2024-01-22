package handler

import (
	"net/http"

	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/labstack/echo/v4"
)

type uvdhandler struct {
	ser *services.Services
}

func Uvdhandlerinit(s *services.Services) *uvdhandler {
	return &uvdhandler{ser: s}
}

func (s *uvdhandler) GetNotification(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	notf, err := s.ser.Notification.GetNotification(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notf)
}
func (s *uvdhandler) CreateNotification(c echo.Context) error {
	u := new(model.Uvd)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := s.ser.Notification.CreateNotification(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
