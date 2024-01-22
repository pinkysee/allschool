package handler

import (
	"net/http"

	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/labstack/echo/v4"
)

type rasphandler struct {
	ser *services.Services
}

func rasphandlerinit(s *services.Services) *rasphandler {
	return &rasphandler{ser: s}
}

func (s *rasphandler) GetRasp(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	raps, err := s.ser.Rasp.GetRasp(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, raps)
}
func (s *rasphandler) SetRasp(c echo.Context) error {
	raps := new(model.Rasp)
	if err := c.Bind(raps); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := s.ser.Rasp.SetRasp(raps)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
func (s *rasphandler) GetAllRasp(c echo.Context) error {
	rasp, err := s.ser.Rasp.GetAllRasp()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rasp)
}
