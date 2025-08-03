package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"main.go/presentations"
	"main.go/repositories"
	"main.go/services"
)

func main() {
	db, err := gorm.Open(
		sqlite.Open("./.data/db.sqlite"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent), TranslateError: true},
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect database"))
	}

	err = db.AutoMigrate(&repositories.Valute{})
	if err != nil {
		panic(errors.Wrap(err, "failed to migrate database"))
	}

	repository := repositories.NewRepository(db)

	service := services.NewService(repository)

	presentation := presentations.NewPresentation(service, repository)

	err = presentation.BuildApp()
	if err != nil {
		log.Error().Stack().Err(err)
	}
}
