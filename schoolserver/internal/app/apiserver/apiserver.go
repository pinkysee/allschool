package apiserver

import (
	"net/http"

	"github.com/PINKYSEE/schoolserver/internal/app/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type apiserver struct {
	config  *Config
	router  *echo.Echo
	handler *handler.Handling
}

func New(config *Config, hand *handler.Handling) *apiserver {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		// Обработка ошибки парсинга уровня логирования
		level = logrus.InfoLevel // Установка уровня логирования по умолчанию
		logrus.Warn("Invalid log level, using default: Info")
	}
	logrus.SetLevel(level)
	return &apiserver{
		config:  config,
		router:  echo.New(),
		handler: hand,
	}
}

func (s *apiserver) Start() error {
	s.configRouter()
	logrus.Info("s handler", s.handler)
	if err := s.router.Start(s.config.BindAddr); err != nil {
		return err
	}

	logrus.Info("Starting API server")
	return nil
}

func (s *apiserver) configRouter() {
	logrus.Info("ff ", s)
	s.router.Static("/docs", "../../internal/app/file/docs")
	s.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://194.87.232.129:3000"},               // Разрешенные источники
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE}, // Разрешенные методы
	}))
	s.router.GET("/hello", hello)
	s.router.POST("/auth", s.handler.Auth)
	s.router.POST("/deluser", s.handler.Delete)
	s.router.GET("/allteacher", s.handler.GetAllTeachers)
	s.router.GET("/allzav", s.handler.GetAllZav)
	s.router.POST("/register", s.handler.Register)
	s.router.GET("/getnews", s.handler.GetNews)
	s.router.POST("/getdocs", s.handler.Getdocs)
	s.router.POST("/deletenews", s.handler.DeleteNews)
	s.router.POST("/addnews", s.handler.CreateNews)
	s.router.POST("/deletedocs", s.handler.DeleteDocs)
	s.router.POST("/adddocs", s.handler.Adddocs) // TODO: мидл на не пустые поля
	s.router.GET("/checkuser", s.handler.CheckJWT)
	s.router.GET("/getalldocs", s.handler.GetAllDocs)
	/* iframe+ */
	s.router.GET("/getiframe", s.handler.GetIframe)
	s.router.GET("/getalliframe", s.handler.GetAllIframe)
	s.router.POST("/addiframe", s.handler.AddIframe)
	s.router.POST("/deleteiframe", s.handler.DeleteIframe)
	/* class+ */
	s.router.GET("/getclass", s.handler.GetClass)
	s.router.POST("/addclass", s.handler.AddClass)
	s.router.POST("/deleteclass", s.handler.DeleteClass)
	/* notf */
	s.router.POST("/setnotf", s.handler.CreateNotification)
	s.router.POST("/getnotf", s.handler.GetNotification)
	/* rasp */
	s.router.POST("/getrasp", s.handler.GetRasp)
	s.router.POST("/setrasp", s.handler.SetRasp)
	s.router.GET("/getallrasp", s.handler.GetAllRasp)
	s.router.POST("/getuserbyclass", s.handler.GetUserByClass)
	s.router.POST("/deleteuserbyclass", s.handler.DeleteUserByClass)
}

func hello(c echo.Context) error {
	logrus.Info("Handling /hello request")
	return c.String(http.StatusOK, "hello")
}
