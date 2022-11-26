package repo

import (
	accountsModel "metroanno-api/app/accounts/domain/models"
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsRepo) CreateTheory(ctx echo.Context, document models.Document) error {
	err := a.MySQL.DB.Table("documents").Create(&document).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AnnotationsRepo) GetDocumentsById(ctx echo.Context, documentId int64) (*models.Document, error) {
	document := &models.Document{}
	err := a.MySQL.DB.Table("documents").Where("id = ?", documentId).First(&document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsRepo) GetAllDocuments(ctx echo.Context) (*[]models.Document, error) {
	document := &[]models.Document{}
	err := a.MySQL.DB.Table("documents").Find(&document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsRepo) UpdateDocumentsById(ctx echo.Context, param request.ReqEditDocument) (*models.Document, error) {
	documentfind := &models.Document{}
	err := a.MySQL.DB.Table("documents").Where("id=?", param.DocumentId).First(documentfind).Error
	if err != nil {
		return nil, err
	}

	document := &models.Document{
		Id:                               param.DocumentId,
		SubjectId:                        param.SubjectId,
		LearningOutcome:                  param.LearningOutcome,
		TextDocument:                     param.TextDocument,
		MinNumberOfQuestionsPerAnnotator: param.MinNumberOfQuestions,
		MinNumberOfAnnotators:            param.MinNumberOfAnnotators,
	}
	err = a.MySQL.DB.Table("documents").Where("id = ?", param.DocumentId).Updates(document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsRepo) GetDocumentIdByUserId(ctx echo.Context) (int64, error) {
	modelAccount := &accountsModel.User{}
	err := a.MySQL.DB.Table(TableUsers).Where("id = ?", ctx.Get("user_id")).Find(&modelAccount).Error
	if err != nil {
		return 0, err
	}
	if modelAccount.CurrentDocumentID == 0 {
		return 0, ErrDocumentNotFound
	}

	modelDocument := &models.Document{}
	err = a.MySQL.DB.Table(TableDocuments).Where("id = ?", modelAccount.CurrentDocumentID).First(&modelDocument).Error
	if err != nil {
		if err.Error() == "record not found" {
			return 0, ErrDocumentIsDeleted
		}
		return 0, err
	}

	return modelAccount.CurrentDocumentID, nil
}

func (a *AnnotationsRepo) DeleteDocumentsByID(ctx echo.Context, id int64) (*models.Document, error) {
	document := &models.Document{
		Id: id,
	}

	documentGet := &models.Document{}
	err := a.MySQL.DB.Table("documents").First(documentGet, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	err = a.MySQL.DB.Table("documents").Delete(&document, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}
