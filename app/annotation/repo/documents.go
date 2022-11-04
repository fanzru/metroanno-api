package repo

import (
	"metroanno-api/app/annotation/domain/models"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsRepo) CreateTheory(ctx echo.Context, document models.Document) error {
	err := a.MySQL.DB.Table("documents").Create(&document).Error
	if err != nil {
		return err
	}
	return nil
}
