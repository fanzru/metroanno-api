package repo

import (
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
	}
	err = a.MySQL.DB.Table("documents").Where("id = ?", param.DocumentId).Updates(document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}
