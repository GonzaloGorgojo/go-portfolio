package registry

import (
	controller "github.com/GonzaloGorgojo/go-backend-portfolio/src/adapters/controllers"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewExperienceController()
}
