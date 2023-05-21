package repo

import (
	"metroanno-api/app/annotation/domain/models"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsRepo) GetAllSubject(ctx echo.Context) ([]models.Subject, error) {
	data := []models.Subject{}
	err := a.MySQL.DB.Table("subjects").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
