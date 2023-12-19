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
	// if !params.IsValidEnum(a.Cfg) {
	// 	return errs.ErrEnumNotValid
	// }

	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return err
	}

	// Set zona waktu ke Jakarta (Waktu Indonesia Barat)
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	// Dapatkan waktu saat ini di zona waktu Jakarta
	currentTime := time.Now().In(location)

	history := models.Histories{
		Name:            fmt.Sprintf(`history %v`, currentTime),
		CreatedAt:       time.Now(),
		DeletedAt:       nil,
		UserID:          userID,
		ReadingMaterial: params.ReadingMaterial,
	}

	err = a.QuestionGenerationRepo.BulkInsertQuestions(ctx, params, history)
	if err != nil {
		return err
	}

	return nil
}

func (a *QuestionGenerationApp) GetHistoryQuestionUser(ctx echo.Context, params params.FilterQuestions) ([]models.Histories, error) {
	questions, err := a.QuestionGenerationRepo.FindQuestions(ctx, params.UserID, params.HistoryID)
	if err != nil {
		return nil, err
	}

	return questions, nil
}
