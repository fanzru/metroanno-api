package repo

import (
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"

	"github.com/labstack/echo/v4"
)

type Impl interface {
	CreateTheory(ctx echo.Context, document models.Document) error
}

type AnnotationsRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(a AnnotationsRepo) AnnotationsRepo {
	return a
}
