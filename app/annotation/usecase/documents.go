package usecase

import (
	"fmt"
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) AddDocument(ctx echo.Context, param request.ReqAddDocument) error {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return err
	}

	err = a.AnnotationsRepo.CreateTheory(ctx, models.Document{
		Id:                                     0,
		SubjectId:                              param.SubjectId,
		LearningOutcome:                        param.LearningOutcome,
		TextDocument:                           param.TextDocument,
		MinNumberOfQuestionsPerAnnotator:       int64(param.MinNumberOfQuestions),
		MinNumberOfAnnotators:                  int64(param.MinNumberOfAnnotators),
		CurrentNumberOfAnnotatorsAssigned:      0,
		CurrentTotalNumberOfQuestionsAnnotated: 0,
		CreatedByUserId:                        userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AnnotationsApp) GetDocumentsById(ctx echo.Context, documentId int64) (*models.Document, error) {
	document, err := a.AnnotationsRepo.GetDocumentsById(ctx, documentId)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (a *AnnotationsApp) GetDocumentsByCreatedBy(ctx echo.Context) ([]models.Document, error) {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return nil, err
	}

	documents, err := a.AnnotationsRepo.GetDocumentsByCreatedBy(ctx, userID)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (a *AnnotationsApp) GetAllDocuments(ctx echo.Context) ([]models.Document, error) {
	documents, err := a.AnnotationsRepo.GetAllDocuments(ctx)
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (a *AnnotationsApp) UpdateDocumentsById(ctx echo.Context, param request.ReqEditDocument) (*models.Document, error) {
	document, err := a.AnnotationsRepo.UpdateDocumentsById(ctx, param)
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsApp) DeleteDocumentsByID(ctx echo.Context, id int64) (*models.Document, error) {
	result, err := a.AnnotationsRepo.DeleteDocumentsByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
