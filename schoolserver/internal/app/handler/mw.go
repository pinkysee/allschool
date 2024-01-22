package handler

import (
	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/labstack/echo/v4"
)

type Mwrep struct {
	rep *repository.Repository
}

func Mwrepinit(r *repository.Repository) *Mwrep {
	return &Mwrep{rep: r}
}

func UserRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
