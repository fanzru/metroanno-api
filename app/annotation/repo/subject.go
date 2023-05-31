package repo

import (
	modelsaccount "metroanno-api/app/accounts/domain/models"
	"metroanno-api/app/annotation/domain/models"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsRepo) GetAllSubject(ctx echo.Context) ([]models.Subject, error) {
	data := []models.Subject{}
	err := a.MySQL.DB.Table("subjects").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AnnotationsRepo) GetAllUsersSubjectWithUserId(ctx echo.Context, userId int64) ([]*modelsaccount.UsersSubjects, error) {
	data := []*modelsaccount.UsersSubjects{}
	err := a.MySQL.DB.Table("users_subjects").Where("user_id = ?", userId).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
