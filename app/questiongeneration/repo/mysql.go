package repo

import (
	"metroanno-api/app/questiongeneration/domain/models"
	"metroanno-api/app/questiongeneration/domain/request"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"
	"time"

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

func (repo *QuestionGenerationRepo) BulkInsertQuestions(ctx echo.Context, params request.ReqSaveQuestions, history models.Histories) error {

	tx := repo.MySQL.DB.Begin()

	err := tx.Create(&history).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	datas := []models.QuestionsHistory{}
	for _, p := range params.SaveQuestions {
		datas = append(datas, models.QuestionsHistory{
			ID:         0,
			Difficulty: p.Difficulty,
			SourceText: p.SourceText,
			Topic:      p.Topic,
			Random:     p.Random,
			Bloom:      p.Bloom,
			Graesser:   p.Graesser,
			CreatedAt:  time.Now(),
			DeletedAt:  nil,
			HistoryID:  history.ID,
			Question:   p.Question,
			Answer:     p.Answer,
			Type:       p.Type,
		})
	}

	if len(datas) == 0 {
		return nil
	}

	// Using gorm clause to ignore duplicate keys
	err = tx.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&datas).Error
	if err != nil {
		return err
	}

	if tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (repo *QuestionGenerationRepo) FindQuestions(ctx echo.Context, userID, questionID int64) ([]models.Histories, error) {
	var histories []models.Histories
	query := repo.MySQL.DB

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	if questionID != 0 {
		query = query.Where("id = ?", questionID)
	}

	err := query.Preload("Questions").Order("created_at DESC").Find(&histories).Error
	if err != nil {
		return nil, err
	}

	return histories, nil
}
