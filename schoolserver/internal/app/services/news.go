package services

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/model"
)

type NewsService struct {
	rep *repository.Repository
}

func NewsServiceinit(repo *repository.Repository) *NewsService {
	return &NewsService{rep: repo}
}
func (s *NewsService) CreateNews(c *model.News, v *multipart.FileHeader) error {
	location, err := time.LoadLocation("Europe/Moscow")
	c.Created_at = time.Now().In(location).Format("2006-01-02 15:04:05")
	var extension string
	c.Preview = randname()
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
	c.Preview = randname() + extension
	ss := s.rep.News.CreateNews(c)
	if ss != nil {
		return ss
	}
	ff := os.MkdirAll("../../internal/app/file/image/", os.ModePerm)
	if ff != nil {
		return ff
	}

	dst, err := os.Create("../../internal/app/file/image/" + c.Preview)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Копируем содержимое файла в созданный файл на сервере
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	defer src.Close()
	return nil
}
func (s *NewsService) GetNews() ([]model.News, error) { return s.rep.GetNews() }
func (s *NewsService) DeleteNews(c *model.News) error {
	err := os.Remove("../../internal/app/file/image/" + c.Preview)
	if err != nil {
		return err
	}
	return s.rep.News.DeleteNews(c.Id)
}
