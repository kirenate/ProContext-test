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
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode   string `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   string `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

func NewRepository(db *gorm.DB) *Repository {
	repository := Repository{db: db}
	return &repository
}

func (r *Repository) SaveValute(valute *Valute) error {
	err := r.db.Create(&valute).Error
	if err != nil {
		return errors.Wrap(err, "failed to save valute")
	}
	return nil
}

func (r *Repository) GetMaxValute() (*Valute, error) {
	var val Valute
	res := r.db.Order("Value desc").Find(&val)
	if res.Error != nil {
		return nil, res.Error
	}

	return &val, nil
}

func (r *Repository) GetMinValute() (*Valute, error) {
	var val Valute
	res := r.db.Order("Value asc").Find(&val)
	if res.Error != nil {
		return nil, res.Error
	}

	return &val, nil
}

func (r *Repository) GetAllRecords() (*[]Valute, error) {
	var vals []Valute
	res := r.db.Select("*").Find(&vals)
	if res.Error != nil {
		return nil, res.Error
	}

	return &vals, nil
}
