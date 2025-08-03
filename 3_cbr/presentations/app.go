package presentations

import (
	"fmt"
	"github.com/pkg/errors"
	"main.go/repositories"
	"main.go/services"
	"strconv"
	"strings"
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

	maxim, err := r.repository.GetMaxValute()
	if err != nil {
		return errors.Wrap(err, "failed to get max valute")
	}

	minim, err := r.repository.GetMinValute()
	if err != nil {
		return errors.Wrap(err, "failed to get min valute")
	}
	all, err := r.repository.GetAllRecords()
	if err != nil {
		return errors.Wrap(err, "failed to get all records")
	}
	avg := 0.0
	for _, v := range *all {

		tmp := strings.Join(strings.Split(v.Value, ","), ".")

		value, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			return errors.Wrap(err, "failed to parse float")
		}

		avg += value

	}
	avg = avg / float64(len(*all))
	fmt.Println(maxim)
	fmt.Println(minim)
	fmt.Println(avg)

	return nil
}
