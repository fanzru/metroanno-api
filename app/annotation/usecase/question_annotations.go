package usecase

import (
	"fmt"
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) BulkInsertQuestionAnnotations(ctx echo.Context, param request.ReqCreateQuestionAnnoatations) (*[]models.QuestionAnnotation, error) {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return nil, err
	}

	documentID, err := a.AnnotationsRepo.GetDocumentIdByUserId(ctx) // hardcode for now before alter table user
	if err != nil {
		return nil, err
	}

	if documentID != param.DocumentId {
		return nil, ErrDifferentDocumentID
	}

	count, err := a.AnnotationsRepo.CountQuestionAnnotationByDocumentID(ctx, documentID) // hardcode for now before alter table user
	if err != nil {
		return nil, err
	}

	var datas []models.QuestionAnnotation
	for i, v := range param.QuestionAnnoatations {
		datas = append(datas, models.QuestionAnnotation{
			Id:                     0,
			UserID:                 userID,
			DocumentID:             param.DocumentId, // hardcode for now before alter table user
			IsLearningOutcomeShown: param.IsLearningOutcomeShown,
			QuestionOrder:          int64(i) + 1 + count,
			QuestionTypeId:         v.QuestionTypeId,
			Keywords:               v.Keywords,
			QuestionText:           v.QuestionText,
			AnswerText:             v.AnswerText,
			TimeDuration:           param.TimeDuration,
			CreatedAt:              time.Now(),
		})
	}

	result, tx, err := a.AnnotationsRepo.BulkInsertQuestionAnnotations(ctx, datas)
	if err != nil {
		return nil, err
	}

	err = a.AnnotationsRepo.CreateDoneDocumentUser(ctx, param.DocumentId, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	documentDoneUser, err := a.AnnotationsRepo.GetAllDocumentDoneUser(ctx, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx, err = a.AnnotationsRepo.UpdateAddDoneDocumentsById(ctx, request.ReqEditDocument{DocumentId: documentID}, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	ID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user, err := a.AnnotationsRepo.GetUserByID(ctx, ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	documents, err := a.AnnotationsRepo.GetAllDocumentsWithWhere(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// remove done document
	newArrDocuments := []models.Document{}
	for _, doc := range documents {
		if doc.DoneNumberOfAnnotators != doc.MinNumberOfAnnotators {
			found := false
			for _, document := range *documentDoneUser {
				if document.DocumentID == doc.Id {
					found = true
					break
				}
			}
			if !found {
				newArrDocuments = append(newArrDocuments, doc)
			}
		}
	}

	if len(newArrDocuments) == 0 {
		_, tx, err = a.AnnotationsRepo.UpdateUsersByIdWithTX(ctx, 0, user.Id, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

	} else {
		sort.SliceStable(newArrDocuments, func(i, j int) bool {
			return newArrDocuments[i].DoneNumberOfAnnotators < newArrDocuments[j].DoneNumberOfAnnotators
		})

		_, tx, err = a.AnnotationsRepo.UpdateUsersByIdWithTX(ctx, newArrDocuments[0].Id, user.Id, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
