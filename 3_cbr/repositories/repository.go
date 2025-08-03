package repositories

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valute  Valute   `xml:"name"`
}

type Valute struct {
	NumCode   int     `xml:"NumCode"`
	CharCode  string  `xml:"CharCode"`
	Nominal   int     `xml:"Nominal"`
	Name      string  `xml:"Name"`
	Value     float64 `xml:"Value"`
	VunitRate float64 `xml:"VunitRate"`
}

func NewRepository(db *gorm.DB) *Repository {
	repository := Repository{db: db}
	return &repository
}

func (r *Repository) SaveValute(valute *Valute) error {
	err := r.db.Save(&valute).Error
	if err != nil {
		return errors.Wrap(err, "failed to save valute")
	}
	return nil
}
