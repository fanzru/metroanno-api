package repo

import (
	"errors"
	"metroanno-api/app/annotation/domain/models"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsRepo) CreateDoneDocumentUser(ctx echo.Context, docId int64, userID int64) error {
	doneDocUser := &models.DoneDocumentUser{
		Id:         0,
		UserID:     userID,
		DocumentID: docId,
		Done:       true,
	}
	err := a.MySQL.DB.Table("done_document_user").Create(doneDocUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AnnotationsRepo) GetDoneDocumentUser(ctx echo.Context, docId int64, userID int64) (bool, error) {
	doneDocUser := []models.DoneDocumentUser{}
	err := a.MySQL.DB.Table("done_document_user").Where("document_id = ? AND user_id = ?", docId, userID).Find(&doneDocUser).Error
	if err != nil {
		return false, err
	}
	if len(doneDocUser) > 0 {
		return false, errors.New("please check your document")
	}
	return true, nil
}
