package services

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	rep *repository.Repository
}

const (
	key = "dsgfdhgdhdsghdfghfdghdrr"
)

type Claims struct {
	jwt.StandardClaims
	Iss string     `json:"iss"`
	Sub model.User `json:"sub"`
	Foo int        `json:"foo"`
	Iat time.Time  `json:"iat"`
}

func UserServiceinit(repo *repository.Repository) *UserService {
	return &UserService{rep: repo}
}
func (s *UserService) GetUserByClass(c *model.User) ([]model.User, error) {
	return s.rep.User.GetUserByClass(c)
}
func (s *UserService) DeleteUserByClass(c *model.User) error {
	return s.rep.User.DeleteUserByClass(c)
}
func (s *UserService) GenerateJWT(c *model.User) (string, error) {

	logrus.Info(c)
	us, err := s.rep.GetUser(c)
	if err != nil && err.Error() != "sql: no rows in result set" {
		logrus.Fatal(err)
	}
	dehashpass := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(c.Password))
	logrus.Info("fffa ", dehashpass, us.Password, c.Password)
	if us.Login == "" && dehashpass != nil {
		return "", nil
	}
	tokeninfo := &Claims{
		jwt.StandardClaims{},
		"schoolserv",
		us,
		rand.Intn(1000000000),
		time.Now(),
	}
	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokeninfo)
	logrus.Info(tokeninfo)
	// Sign the token
	sd, err := token.SignedString([]byte(key))
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return sd, nil
}

func (s *UserService) Register(c *model.User) (int, error, string) {
	pas := GeneratePassword()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pas), bcrypt.DefaultCost)
	log.Info(pas)
	if err != nil {
		return 0, err, ""
	}
	c.Password = string(hashedPassword)
	c.Avatar = rand.Intn(32)
	id, err := s.rep.User.AddUser(c)
	if err != nil {
		return 0, err, ""
	}
	return id, err, pas
}
func GeneratePassword() string {
	var password string
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!№;%:?*()_+="
	f := strings.Split(s, "")
	for i := 0; i < 8; i++ {
		password = password + f[rand.Intn(len(f)-1)]
	}
	return password
}
func (s *UserService) Delete(userID int) (string, error) {
	return s.rep.User.DeleteUser(userID)
}

func (s *UserService) GetAllTeacher() ([]model.User, error) {
	return s.rep.User.GetAllTeacher()
}

func (s *UserService) GetAllZav() ([]model.User, error) {
	return s.rep.User.GetAllZav()
}
func (s *UserService) CheckJWT(jw string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(jw, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt method")
		}

		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	data, ok := token.Claims.(*Claims)
	u := new(model.User)
	if !ok {
		return nil, errors.New("token claims are not of type claims")
	}
	u.Name = data.Sub.Name
	u.Avatar = data.Sub.Avatar
	u.Role = data.Sub.Role
	return u, nil
}
