package usecase

import (
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) CreateQuestionTypes(ctx echo.Context, param request.ReqAddQuestionType) (*models.QuestionType, error) {
	result, err := a.AnnotationsRepo.CreateQuestionTypes(ctx, param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *AnnotationsApp) DeleteQuestionTypes(ctx echo.Context, id int64) (*models.QuestionType, error) {
	result, err := a.AnnotationsRepo.DeleteQuestionTypes(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *AnnotationsApp) GetAllQuestionTypes(ctx echo.Context) (*[]models.QuestionType, error) {
	result, err := a.AnnotationsRepo.GetAllQuestionTypes(ctx)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
