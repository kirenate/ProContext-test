package repositories

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	repository := Repository{db: db}
	return &repository
}
