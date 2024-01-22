package handler

import (
	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/labstack/echo/v4"
)

type Testhandler struct {
	ser *services.Services
}

func Testhandlerinit(s *services.Services) *Userhandler {
	return &Userhandler{ser: s}
}

func (s *Userhandler) Settest(c *echo.Context) error       { return nil }
func (s *Userhandler) Gettest(c *echo.Context) error       { return nil }
func (s *Userhandler) Setanswertest(c *echo.Context) error { return nil }
