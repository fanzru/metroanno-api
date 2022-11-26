package usecase

import (
	"fmt"
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) CreateFeeback(ctx echo.Context, param request.ReqCreateFeedback) (*models.Feedback, error) {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := a.AnnotationsRepo.CreateFeedback(ctx, models.Feedback{
		UserID:       userID,
		DocumentID:   param.DocumentID,
		FeedbackText: param.FeedbackText,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
