package controller

import (
	"net/http"

	model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"
	"github.com/GonzaloGorgojo/go-backend-portfolio/src/usecases/interactor"
	"github.com/labstack/echo"
)

type experienceController struct {
	experienceInteractor interactor.ExperienceInteractor
}

type ExperienceController interface {
	GetExperience(c echo.Context) error
	AddExperience(c echo.Context) error
}

func NewExperienceController(us interactor.ExperienceInteractor) ExperienceController {
	return &experienceController{us}
}

func (uc *experienceController) GetExperience(c echo.Context) error {
	var u []*model.Experience

	u, err := uc.experienceInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *experienceController) AddExperience(c echo.Context) (err error) {
	u := new(model.Experience)

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}

	u, e := uc.experienceInteractor.AddOne(u)

	if e != nil {
		return e
	}

	return c.JSON(http.StatusCreated, u)
}
