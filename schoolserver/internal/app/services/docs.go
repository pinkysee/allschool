package services

import (
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"strings"

	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/model"
)

type DocsRepository struct {
	rep *repository.Repository
}

func DocsRepositoryinit(rep *repository.Repository) *DocsRepository {
	return &DocsRepository{rep: rep}
}

func (s *DocsRepository) GetDocs(page string) ([]model.Docs, error) {
	return s.rep.Docs.GetDocs(page)
}
func (s *DocsRepository) GetAllDocs() ([]model.Docs, error) {
	return s.rep.Docs.GetAllDocs()
}
func (s *DocsRepository) AddDocs(c *model.Docs, v *multipart.FileHeader) (string, error) {
	var extension string
	src, err := v.Open()
	defer src.Close()
	lastDotIndex := strings.LastIndex(v.Filename, ".")

	// Проверяем, что точка найдена и не является последним символом в строке
	if lastDotIndex != -1 && lastDotIndex < len(v.Filename)-1 {
		extension = v.Filename[lastDotIndex:]
	} else {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	c.Path = randname() + extension
	ff := os.MkdirAll("../../internal/app/file/docs/", os.ModePerm)
	if ff != nil {
		return "", ff
	}

	dst, err := os.Create("../../internal/app/file/docs/" + c.Path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Копируем содержимое файла в созданный файл на сервере
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}
	defer src.Close()
	return s.rep.Docs.AddDocs(c)
}

func (s *DocsRepository) DeleteDocs(path string) (string, error) {
	ss, errs := s.rep.Docs.DeleteDocs(path)
	if errs != nil {
		return "", errs
	}
	err := os.Remove("../../internal/app/file/docs/" + path)
	if err != nil {
		return "", err
	}
	return ss, nil
}
func randname() string {
	var name string
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	chars := strings.Split(s, "") // Переместите эту строку за пределы цикла

	for i := 0; i < 64; i++ {
		name = name + chars[rand.Intn(len(chars))]
	}
	return name
}
