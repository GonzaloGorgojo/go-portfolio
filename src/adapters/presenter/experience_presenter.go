package presenter

import model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"

type experiencePresenter struct {
}

type ExperiencePresenter interface {
	ResponseExperience(us []*model.Experience) []*model.Experience
	ResponseAddOne(us *model.Experience) *model.Experience
}

func NewExperiencePresenter() ExperiencePresenter {
	return &experiencePresenter{}
}

func (up *experiencePresenter) ResponseExperience(us []*model.Experience) []*model.Experience {

	return us
}

func (up *experiencePresenter) ResponseAddOne(us *model.Experience) *model.Experience {

	return us
}
