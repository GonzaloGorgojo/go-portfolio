package repository

import (
	model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"
)

type ExperienceRepository interface {
	FindAll(u []*model.Experience) ([]*model.Experience, error)
	AddOne(u *model.Experience) (*model.Experience, error)
}
