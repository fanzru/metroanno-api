package usecase

import (
	"metroanno-api/app/annotation/domain/models"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) GetAllSubject(ctx echo.Context) ([]models.Subject, error) {
	subjects, err := a.AnnotationsRepo.GetAllSubject(ctx)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}
