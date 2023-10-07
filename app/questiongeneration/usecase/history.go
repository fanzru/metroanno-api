package usecase

import (
	"fmt"
	"metroanno-api/app/questiongeneration/domain/models"
	"metroanno-api/app/questiongeneration/domain/params"
	"metroanno-api/app/questiongeneration/domain/request"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *QuestionGenerationApp) SaveQuestions(ctx echo.Context, params request.ReqSaveQuestions) error {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return err
	}

	datas := []models.QuestionsHistory{}
	for _, p := range params.SaveQuestions {
		datas = append(datas, models.QuestionsHistory{
			ID:              0,
			Difficulty:      p.Difficulty,
			ReadingMaterial: p.ReadingMaterial,
			Topic:           p.Topic,
			Random:          p.Random,
			Bloom:           p.Bloom,
			Graesser:        p.Graesser,
			CreatedAt:       time.Now(),
			DeletedAt:       nil,
			UserID:          userID,
		})
	}

	err = a.QuestionGenerationRepo.BulkInsertQuestions(ctx, datas)
	if err != nil {
		return err
	}

	return nil
}

func (a *QuestionGenerationApp) GetHistoryQuestionUser(ctx echo.Context, params params.FilterQuestions) ([]models.QuestionsHistory, error) {
	questions, err := a.QuestionGenerationRepo.FindQuestions(ctx, params.UserID, params.QuestionID)
	if err != nil {
		return nil, err
	}

	return questions, nil
}
