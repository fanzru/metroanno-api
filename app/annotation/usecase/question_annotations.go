package usecase

import (
	"fmt"
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) BulkInsertQuestionAnnotations(ctx echo.Context, param request.ReqCreateQuestionAnnoatations) (*[]models.QuestionAnnotation, error) {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return nil, err
	}

	documentID, err := a.AnnotationsRepo.GetDocumentIdByUserId(ctx) // hardcode for now before alter table user
	if err != nil {
		return nil, err
	}

	count, err := a.AnnotationsRepo.CountQuestionAnnotationByDocumentID(ctx, documentID) // hardcode for now before alter table user
	if err != nil {
		return nil, err
	}

	var datas []models.QuestionAnnotation
	for i, v := range param.QuestionAnnoatations {
		datas = append(datas, models.QuestionAnnotation{
			Id:                     0,
			UserID:                 userID,
			DocumentID:             documentID, // hardcode for now before alter table user
			IsLearningOutcomeShown: param.IsLearningOutcomeShown,
			QuestionOrder:          int64(i) + 1 + count,
			QuestionTypeId:         v.QuestionTypeId,
			Keywords:               v.Keywords,
			QuestionText:           v.QuestionText,
			AnswerText:             v.AnswerText,
			TimeDuration:           param.TimeDuration,
			CreatedAt:              time.Now(),
		})
	}

	result, err := a.AnnotationsRepo.BulkInsertQuestionAnnotations(ctx, datas)
	if err != nil {
		return nil, err
	}

	return result, nil
}
