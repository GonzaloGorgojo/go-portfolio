package registry

import (
	controller "github.com/GonzaloGorgojo/go-backend-portfolio/src/adapters/controllers"
	ap "github.com/GonzaloGorgojo/go-backend-portfolio/src/adapters/presenter"
	ar "github.com/GonzaloGorgojo/go-backend-portfolio/src/adapters/repository"
	"github.com/GonzaloGorgojo/go-backend-portfolio/src/usecases/interactor"
	up "github.com/GonzaloGorgojo/go-backend-portfolio/src/usecases/presenter"
	ur "github.com/GonzaloGorgojo/go-backend-portfolio/src/usecases/repository"
)

func (r *registry) NewExperienceController() controller.ExperienceController {
	return controller.NewExperienceController(r.NewExperienceInteractor())
}

func (r *registry) NewExperienceInteractor() interactor.ExperienceInteractor {
	return interactor.NewExperienceInteractor(r.NewExperienceRepository(), r.NewExperiencePresenter())
}

func (r *registry) NewExperienceRepository() ur.ExperienceRepository {
	return ar.NewExperienceRepository(r.db)
}

func (r *registry) NewExperiencePresenter() up.ExperiencePresenter {
	return ap.NewExperiencePresenter()
}
