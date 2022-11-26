package repo

import (
	"metroanno-api/app/annotation/domain/models"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsRepo) CreateFeedback(ctx echo.Context, feedback models.Feedback) (*models.Feedback, error) {
	models := &models.Feedback{
		Id:           0,
		UserID:       feedback.UserID,
		DocumentID:   feedback.DocumentID,
		FeedbackText: feedback.FeedbackText,
		CreatedAt:    time.Now(),
	}
	err := a.MySQL.DB.Table("feedbacks").Create(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}
