package presenter

import (
	model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"
)

type ExperiencePresenter interface {
	ResponseExperience(u []*model.Experience) []*model.Experience
	ResponseAddOne(u *model.Experience) *model.Experience
}
