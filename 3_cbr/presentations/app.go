package presentations

import (
	"github.com/pkg/errors"
	"main.go/repositories"
	"main.go/services"
)

type Presentation struct {
	Service    *services.Service
	Repository *repositories.Repository
}

func NewPresentation(service *services.Service, repository *repositories.Repository) *Presentation {
	return &Presentation{Service: service, Repository: repository}
}

func (r *Presentation) BuildApp() error {
	err := r.Service.ProcessData()
	if err != nil {
		return errors.Wrap(err, "failed to process data")
	}
	return nil
}
