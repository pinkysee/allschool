package services

import (
	"strings"

	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/model"
)

type DaysRasp struct {
	Monday   []string `json:"Monday" xml:"Monday" form:"Monday" query:"Monday"`
	Tuesday  []string `json:"Tuesday" xml:"Tuesday" form:"Tuesday" query:"Tuesday"`
	Wendsday []string `json:"Wendsday" xml:"Wendsday" form:"Wendsday" query:"Wendsday"`
	Thursday []string `json:"Thursday" xml:"Thursday" form:"Thursday" query:"Thursday"`
	Friday   []string `json:"Friday" xml:"Friday" form:"Friday" query:"Friday"`
}
type AllRasp struct {
	Class string    `json:"class" xml:"class" form:"class" query:"class"`
	Rasp  *DaysRasp `json:"Raspclass" xml:"Raspclass" form:"Raspclass" query:"Raspclass"`
}
type RaspRepository struct {
	rep *repository.Repository
}

func RaspRepositoryinit(rep *repository.Repository) *RaspRepository {
	return &RaspRepository{rep: rep}
}

func (s *RaspRepository) GetRasp(c *model.User) (*DaysRasp, error) {
	rasp, err := s.rep.Rasp.GetRasp(c)
	if err != nil {
		return nil, err
	}
	var formatrasp *DaysRasp
	formatrasp = &DaysRasp{
		Monday:   strings.Split(strings.Split(rasp[0].Lessons, ";")[0], ","),
		Tuesday:  strings.Split(strings.Split(rasp[0].Lessons, ";")[1], ","),
		Wendsday: strings.Split(strings.Split(rasp[0].Lessons, ";")[2], ","),
		Thursday: strings.Split(strings.Split(rasp[0].Lessons, ";")[3], ","),
		Friday:   strings.Split(strings.Split(rasp[0].Lessons, ";")[4], ","),
	}
	return formatrasp, nil
}
func (s *RaspRepository) SetRasp(c *model.Rasp) error {
	return s.rep.Rasp.SetRasp(c)
}
func (s *RaspRepository) GetAllRasp() ([]*AllRasp, error) {
	rasp, err := s.rep.Rasp.GetAllRasp()
	if err != nil {
		return nil, err
	}
	var formatrasp []*AllRasp
	for _, value := range rasp {
		formatrasp = append(formatrasp, &AllRasp{
			Class: value.Classname,
			Rasp: &DaysRasp{
				Monday:   strings.Split(strings.Split(value.Lessons, ";")[0], ","),
				Tuesday:  strings.Split(strings.Split(value.Lessons, ";")[1], ","),
				Wendsday: strings.Split(strings.Split(value.Lessons, ";")[2], ","),
				Thursday: strings.Split(strings.Split(value.Lessons, ";")[3], ","),
				Friday:   strings.Split(strings.Split(value.Lessons, ";")[4], ","),
			},
		})
	}
	return formatrasp, nil
}
