package presentations

import (
	"github.com/pkg/errors"
	"main.go/repositories"
	"main.go/services"
)

type Presentation struct {
	service    *services.Service
	repository *repositories.Repository
}

func NewPresentation(service *services.Service, repository *repositories.Repository) *Presentation {
	return &Presentation{service: service, repository: repository}
}

func (r *Presentation) BuildApp() error {
	err := r.service.ProcessData()
	if err != nil {
		return errors.Wrap(err, "failed to process data")
	}
	err = r.service.GetNeededInfo()
	if err != nil {
		return errors.Wrap(err, "failed to get needed info")
	}
	return nil
}
