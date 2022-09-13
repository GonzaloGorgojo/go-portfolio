package presenter

import model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"

type experiencePresenter struct {
}

type ExperiencePresenter interface {
	ResponseExperience(us []*model.Experience) []*model.Experience
}

func NewExperiencePresenter() ExperiencePresenter {
	return &experiencePresenter{}
}

func (up *experiencePresenter) ResponseExperience(us []*model.Experience) []*model.Experience {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
