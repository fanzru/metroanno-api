package repo

import (
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v9"
)

func (a *AnnotationsRepo) GetAllQuestionTypes(ctx echo.Context) ([]models.QuestionType, error) {
	questionTypes := []models.QuestionType{}
	err := a.MySQL.DB.Table("question_types").Find(&questionTypes).Error
	if err != nil {
		return nil, err
	}
	return questionTypes, nil
}

func (a *AnnotationsRepo) CreateQuestionTypes(ctx echo.Context, param request.ReqAddQuestionType) (*models.QuestionType, error) {
	questionTypes := &models.QuestionType{
		QuestionType: param.QuestionType,
		Description:  param.Description,
		CreatedAt:    time.Now(),
	}
	err := a.MySQL.DB.Table("question_types").Create(&questionTypes).Error
	if err != nil {
		return nil, err
	}
	return questionTypes, nil
}

func (a *AnnotationsRepo) DeleteQuestionTypes(ctx echo.Context, id int64) (*models.QuestionType, error) {
	time := null.Time{
		Valid: true,
		Time:  time.Now(),
	}
	questionTypes := &models.QuestionType{
		Id:        id,
		DeletedAt: time,
	}

	err := a.MySQL.DB.Table("question_types").Delete(&questionTypes, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return questionTypes, nil
}
