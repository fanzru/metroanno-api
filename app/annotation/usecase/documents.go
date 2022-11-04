package usecase

import (
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) AddDocument(ctx echo.Context, param request.ReqAddDocument) error {
	err := a.AnnotationsRepo.CreateTheory(ctx, models.Document{
		Id:                                     0,
		SubjectId:                              param.SubjectId,
		LearningOutcome:                        param.LearningOutcome,
		TextDocument:                           param.TextDocument,
		MinNumberOfQuestionsPerAnnotator:       int64(param.MinNumberOfQuestions),
		MinNumberOfAnnotators:                  0,
		CurrentNumberOfAnnotatorsAssigned:      0,
		CurrentTotalNumberOfQuestionsAnnotated: 0,
	})
	if err != nil {
		return err
	}
	return nil
}
