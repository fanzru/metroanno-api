package repo

import (
	"metroanno-api/app/questiongeneration/domain/models"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type Impl interface {
	BulkInsertQuestions(ctx echo.Context, questions []models.QuestionsHistory) error
	FindQuestions(ctx echo.Context, userID, questionID int64) ([]models.QuestionsHistory, error)
}

type QuestionGenerationRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(a QuestionGenerationRepo) QuestionGenerationRepo {
	return a
}

func (repo *QuestionGenerationRepo) BulkInsertQuestions(ctx echo.Context, questions []models.QuestionsHistory) error {
	if len(questions) == 0 {
		return nil
	}

	// Using gorm clause to ignore duplicate keys
	err := repo.MySQL.DB.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&questions).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *QuestionGenerationRepo) FindQuestions(ctx echo.Context, userID, questionID int64) ([]models.QuestionsHistory, error) {
	var questions []models.QuestionsHistory
	query := repo.MySQL.DB

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	if questionID != 0 {
		query = query.Where("id = ?", questionID)
	}

	err := query.Find(&questions).Error
	if err != nil {
		return nil, err
	}

	return questions, nil
}
