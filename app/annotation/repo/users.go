package repo

import (
	"metroanno-api/app/accounts/domain/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (i *AnnotationsRepo) GetUserByID(ctx echo.Context, ID int64) (models.User, error) {
	var user models.User
	result := i.MySQL.DB.Table("users").Where("id = ?", ID).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, ErrInstanceNotFound
		}
		return user, result.Error
	}
	if result.RowsAffected < 1 {
		return user, ErrInstanceNotFound
	}
	return user, nil
}

func (a *AnnotationsRepo) UpdateUsersById(ctx echo.Context, docId int64, userID int64) (*models.User, error) {
	documentfind := &models.User{}
	err := a.MySQL.DB.Table("users").Where("id=?", userID).First(documentfind).Error
	if err != nil {
		return nil, err
	}

	user := &models.User{
		CurrentDocumentID: docId,
	}
	err = a.MySQL.DB.Table("users").Where("id = ?", userID).Updates(map[string]interface{}{"current_document_id": docId}).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AnnotationsRepo) UpdateUsersByIdWithTX(ctx echo.Context, docId int64, userID int64, tx *gorm.DB) (*models.User, *gorm.DB, error) {
	documentfind := &models.User{}
	err := tx.Table("users").Where("id=?", userID).First(documentfind).Error
	if err != nil {
		return nil, nil, err
	}

	user := &models.User{
		CurrentDocumentID: docId,
	}
	err = tx.Table("users").Where("id = ?", userID).Updates(map[string]interface{}{"current_document_id": docId}).Error
	if err != nil {
		return nil, nil, err
	}
	return user, tx, nil
}
