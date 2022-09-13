package repository

import (
	model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"
	"gorm.io/gorm"
)

type experienceRepository struct {
	db *gorm.DB
}

type ExperienceRepository interface {
	FindAll(u []*model.Experience) ([]*model.Experience, error)
}

func NewExperienceRepository(db *gorm.DB) ExperienceRepository {
	return &experienceRepository{db}
}

func (ur *experienceRepository) FindAll(u []*model.Experience) ([]*model.Experience, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
