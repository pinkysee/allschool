package handler

import (
	"net/http"

	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Api struct {
	Token string
}
type Userhandler struct {
	ser *services.Services
}

func Userhandlerinit(s *services.Services) *Userhandler {
	return &Userhandler{ser: s}
}
func (s *Userhandler) GetUserByClass(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	us, err := s.ser.User.GetUserByClass(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, us)
}
func (s *Userhandler) DeleteUserByClass(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := s.ser.User.DeleteUserByClass(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
func (s *Userhandler) Auth(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	d := &model.User{Login: u.Login, Password: u.Password}
	token, err := s.ser.GenerateJWT(d)
	if err != nil {
		logrus.Error(err)
	} else if token == "" {
		jsontoken := &Api{Token: "UnRegistered"}
		return c.JSON(http.StatusUnauthorized, jsontoken)
	}
	jsontoken := &Api{Token: token}
	return c.JSON(http.StatusOK, jsontoken)
}
func (s *Userhandler) Register(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	_, _, pas := s.ser.User.Register(u)
	u.Password = pas
	return c.JSON(http.StatusCreated, u)
}
func (s *Userhandler) Delete(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	ans, err := s.ser.User.Delete(int(u.ID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ans)
}
func (s *Userhandler) GetAllTeachers(c echo.Context) error {
	ans, err := s.ser.User.GetAllTeacher()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ans)
}
func (s *Userhandler) GetAllZav(c echo.Context) error {
	ans, err := s.ser.User.GetAllZav()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ans)
}
func (s *Userhandler) CheckJWT(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	ans, err := s.ser.User.CheckJWT(header)
	if err != nil && ans == nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusAccepted, ans)
}
