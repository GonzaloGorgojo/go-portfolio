package interactor

import (
	model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"
	"github.com/GonzaloGorgojo/go-backend-portfolio/src/usecases/presenter"
	"github.com/GonzaloGorgojo/go-backend-portfolio/src/usecases/repository"
)

type experienceInteractor struct {
	ExperienceRepository repository.ExperienceRepository
	ExperiencePresenter  presenter.ExperiencePresenter
}

type ExperienceInteractor interface {
	Get(u []*model.Experience) ([]*model.Experience, error)
}

func NewExperienceInteractor(r repository.ExperienceRepository, p presenter.ExperiencePresenter) ExperienceInteractor {
	return &experienceInteractor{r, p}
}

func (us *experienceInteractor) Get(u []*model.Experience) ([]*model.Experience, error) {
	u, err := us.ExperienceRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.ExperiencePresenter.ResponseExperience(u), nil
}
