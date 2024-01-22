package services

import (
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/sirupsen/logrus"
)

type IframeServices struct {
	rep *repository.Repository
}

func IframeServicesinit(repo *repository.Repository) *IframeServices {
	return &IframeServices{rep: repo}
}
func (s *IframeServices) AddIframe(c *model.Iframe, v *multipart.FileHeader) error {
	var extension string
	src, err := v.Open()
	defer src.Close()
	lastDotIndex := strings.LastIndex(v.Filename, ".")

	// Проверяем, что точка найдена и не является последним символом в строке
	if lastDotIndex != -1 && lastDotIndex < len(v.Filename)-1 {
		extension = v.Filename[lastDotIndex:]
	} else {
		return nil
	}
	if err != nil {
		return err
	}
	c.Path = randname() + extension
	ff := os.MkdirAll("../../internal/app/file/iframe/", os.ModePerm)
	if ff != nil {
		logrus.Error(err)
		return ff
	}

	dst, err := os.Create("../../internal/app/file/iframe/" + c.Path)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer dst.Close()

	// Копируем содержимое файла в созданный файл на сервере
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	defer src.Close()
	return s.rep.Iframe.AddIframe(c)
}

func (s *IframeServices) GetIframe(page string) ([]model.Iframe, error) {
	return s.rep.GetIframe(page)
}
func (s *IframeServices) DeleteIframe(c *model.Iframe) error {
	errs := s.rep.Iframe.DeleteIframe(c)
	if errs != nil {
		return errs
	}
	err := os.Remove("../../internal/app/file/iframe/" + c.Path)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
func (s *IframeServices) GetAllIframe() ([]model.Iframe, error) {
	return s.rep.GetAllIframe()
}
