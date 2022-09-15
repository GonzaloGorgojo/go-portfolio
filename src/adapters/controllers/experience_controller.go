package controller

import (
	"net/http"

	model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"
	"github.com/GonzaloGorgojo/go-backend-portfolio/src/usecases/interactor"
)

type experienceController struct {
	experienceInteractor interactor.ExperienceInteractor
}

type ExperienceController interface {
	GetExperience(c Context) error
	AddExperience(c Context) error
}

func NewExperienceController(us interactor.ExperienceInteractor) ExperienceController {
	return &experienceController{us}
}

func (uc *experienceController) GetExperience(c Context) error {
	var u []*model.Experience

	u, err := uc.experienceInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *experienceController) AddExperience(c Context) error {
	var u *model.Experience

	u, err := uc.experienceInteractor.AddOne(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
